package service

import (
	"fmt"
	"scrim-api/database"
	"scrim-api/model"
	"strconv"
)

func TeamCreate(userId string, data model.TeamCreateReq) (*model.TeamCreateResp, error) {
	var lastTeamId int
	err := database.Db.QueryRow("SELECT team_id FROM \"team\" ORDER BY team_id DESC LIMIT 1").Scan(&lastTeamId)
	if err != nil {
		return nil, err
	}

	fmt.Printf("last team_id: %v", lastTeamId)

	// find game_id
	var gameId int

	err = database.Db.QueryRow("SELECT game_id FROM \"user\" WHERE user_id = $1", userId).Scan(&gameId)
	if err != nil {
		return nil, err
	}

	if data.TeamLogo != nil {
		_, err = database.Db.Exec("INSERT INTO \"team\" (team_id, team_name, team_logo, game_id) VALUES ($1, $2, $3, $4)",
			lastTeamId+1, data.TeamName, data.TeamLogo, gameId)

		if err != nil {
			return nil, err
		}
	} else {
		_, err = database.Db.Exec("INSERT INTO \"team\" (team_id, team_name, game_id) VALUES ($1, $2, $3)",
			lastTeamId+1, data.TeamName, gameId)

		if err != nil {
			return nil, err
		}
	}

	lastTeamId = lastTeamId + 1

	sql_statement := "UPDATE \"user\" SET team_id = $2, role = $3 WHERE user_id = $1;"
	_, err = database.Db.Exec(sql_statement, userId, lastTeamId, "Manager")
	if err != nil {
		return nil, err
	}

	return &model.TeamCreateResp{
		TeamId:   lastTeamId,
		TeamName: data.TeamName,
		TeamLogo: data.TeamLogo,
	}, nil
}

func TeamUpdate(data model.TeamUpdate) error {
	var err error
	if data.TeamLogo != nil {
		_, err = database.Db.Exec("UPDATE \"team\" SET team_name = $1, team_logo = $2 WHERE team_id = $3",
			data.TeamName,
			data.TeamLogo,
			data.TeamId)

		if err != nil {
			return err
		}
	} else {
		_, err = database.Db.Exec("UPDATE \"team\" SET team_name = $1  WHERE team_id = $2",
			data.TeamName,
			data.TeamId)
		if err != nil {
			return err
		}
	}

	return nil
}

func TeamJoin(data model.TeamJoin) (*model.TeamDetail, error) {
	var teamID int
	err := database.Db.QueryRow("SELECT team_id FROM team WHERE invite_code = $1", data.InviteCode).Scan(&teamID)
	if err != nil {
		fmt.Println("Error querying team table:", err)
		return nil, err
	}

	_, err = database.Db.Exec("UPDATE \"user\" SET team_id = $1 WHERE user_id = $2", teamID, data.UserId)
	fmt.Println("Error updating user table:", err)
	if err != nil {
		return nil, err
	}

	teamDetail, err := TeamDetailGet(strconv.Itoa(teamID))
	if err != nil {
		return nil, err
	}

	return teamDetail, nil
}

func TeamMemberGet(teamId string) (*model.TeamMember, error) {
	var teamMembers model.TeamMember
	rows, err := database.Db.Query("SELECT user_id, nickname, role FROM \"user\" WHERE team_id = $1", teamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member model.Member
		err := rows.Scan(&member.UserId, &member.Nickname, &member.Role)
		if err != nil {
			return nil, err
		}
		teamMembers.Members = append(teamMembers.Members, member)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &teamMembers, nil
}

func TeamDetailGet(teamId string) (*model.TeamDetail, error) {
	var teamDetail model.TeamDetail
	rows, err := database.Db.Query("SELECT user_id, nickname, role FROM \"user\" WHERE team_id = $1", teamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member model.Member
		err := rows.Scan(&member.UserId, &member.Nickname, &member.Role)
		if err != nil {
			return nil, err
		}
		teamDetail.Members = append(teamDetail.Members, member)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	err = database.Db.QueryRow("SELECT team_name, game_id, team_logo, invite_code, invite_flag FROM team WHERE team_id = $1", teamId).Scan(&teamDetail.TeamName, &teamDetail.GameId, &teamDetail.TeamLogo, &teamDetail.InviteCode, &teamDetail.InviteFlag)
	if err != nil {
		return nil, err
	}

	return &teamDetail, nil
}

func TeamSetFlagInviteCode(teamId string) (*model.TeamInviteCodeGet, error) {
	_, err := database.Db.Exec("UPDATE \"team\" SET invite_flag = $1 WHERE team_id = $2",
		true, teamId)

	if err != nil {
		return nil, err
	}

	var inviteCode string
	err = database.Db.QueryRow("SELECT invite_code FROM team WHERE team_id = $1", teamId).Scan(&inviteCode)
	if err != nil {
		return nil, err
	}

	return &model.TeamInviteCodeGet{
		InviteCode: inviteCode,
	}, nil
}
