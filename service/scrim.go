package service

import (
	"database/sql"
	"fmt"
	"scrim-api/database"
	"scrim-api/model"
	"strconv"
)

func ScrimPost(data model.ScrimPost) (*int, error) {
	// update an existing scrim
	var detail model.ScrimDetail
	err := database.Db.QueryRow("SELECT scrim_id FROM scrim WHERE scrim_date = $1 and scrim_time = $2 and team_id = $3", data.ScrimDate, data.ScrimTime, data.TeamId).Scan(&detail.ScrimId)
	if err != nil {
		return nil, err
	}

	// create a new scrim
	var lastScrimId int
	err = database.Db.QueryRow("SELECT scrim_id FROM scrim ORDER BY scrim_id DESC LIMIT 1").Scan(&lastScrimId)
	if err != nil {
		return nil, err
	}

	fmt.Printf("last scrim_id: %v", lastScrimId)

	if detail.ScrimId == 0 {
		_, err = database.Db.Exec("INSERT INTO \"scrim\" (scrim_id, scrim_date, scrim_time, scrim_map, game_id, team_id) VALUES ($1, $2, $3, $4, $5, $6)",
			lastScrimId+1, data.ScrimDate, data.ScrimTime, data.ScrimMap, data.GameId, data.TeamId)
		if err != nil {
			return nil, err
		}
		lastScrimId = lastScrimId + 1
	} else {
		_, err = database.Db.Exec("UPDATE \"scrim\" SET scrim_map = $1 WHERE scrim_id = $2;", data.ScrimMap, detail.ScrimId)
		if err != nil {
			return nil, err
		}
	}

	return &lastScrimId, nil
}

func ScrimMakeOffer(data model.ScrimMakeOffer) error {
	_, err := database.Db.Exec("INSERT INTO \"scrim_offer\" (scrim_id, team_id) VALUES ($1, $2)",
		data.ScrimId, data.TeamId)

	if err != nil {
		return err
	}

	return nil
}

func ScrimAcceptOffer(data model.ScrimAcceptOffer) error {
	_, err := database.Db.Exec("UPDATE \"scrim_offer\" SET offer_status = $1 WHERE scrim_id = $2 AND team_id = $3;",
		"accepted", data.ScrimId, data.TeamId)

	if err != nil {
		return err
	}

	_, err = database.Db.Exec("UPDATE \"scrim\" SET offer_team_id = $1, scrim_status = $3 WHERE scrim_id = $2;",
		data.TeamId, data.ScrimId, "matched")
	if err != nil {
		return err
	}

	return nil
}

func ScrimCancelMatch(data model.ScrimCancelMatch) error {
	_, err := database.Db.Exec("UPDATE \"scrim\" SET scrim_status = $1, offer_team_id = $3 WHERE scrim_id = $2;",
		"unmatched", data.ScrimId, nil)
	if err != nil {
		return err
	}

	_, err = database.Db.Exec("DELETE FROM \"scrim_offer\"  WHERE scrim_id = $1 AND team_id = $2;",
		data.ScrimId, data.TeamId)
	if err != nil {
		return err
	}

	return nil
}

func ScrimDelete(data model.ScrimDelete) error {
	_, err := database.Db.Exec("DELETE FROM \"scrim\"  WHERE scrim_id = $1;",
		data.ScrimId)
	if err != nil {
		return err
	}

	return nil
}

func ScrimGetOffer(teamId string) (*model.ScrimGet, error) {
	var scrims model.ScrimGet
	var rows *sql.Rows
	var err error

	rows, err = database.Db.Query("SELECT scrim.scrim_id, scrim_offer.team_id, t.team_logo, t.team_name, scrim.scrim_map, scrim.scrim_date, scrim.scrim_time FROM scrim INNER JOIN scrim_offer ON scrim.scrim_id = scrim_offer.scrim_id INNER JOIN team AS t ON scrim_offer.team_id = t.team_id WHERE scrim.team_id = $1 AND (scrim.scrim_date > CURRENT_DATE OR (scrim.scrim_date = CURRENT_DATE  AND scrim.scrim_time > CURRENT_TIME )) ORDER BY scrim.scrim_date ASC, scrim.scrim_time ASC;", teamId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var detail model.ScrimDetail
		err := rows.Scan(&detail.ScrimId, &detail.TeamId, &detail.TeamLogo, &detail.TeamName, &detail.ScrimMap, &detail.ScrimDate, &detail.ScrimTime)
		if err != nil {
			return nil, err
		}
		scrims.Scrims = append(scrims.Scrims, detail)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &scrims, nil
}

func ScrimGet(data model.ScrimGetReq) (*model.ScrimQueryResp, error) {
	var scrims model.ScrimQueryResp
	var rows *sql.Rows
	var err error
	if data.ScrimMap != nil {
		rows, err = database.Db.Query("SELECT scrim.scrim_id, scrim.team_id, team.team_logo, team.team_name, scrim.scrim_map, scrim.scrim_date, scrim.scrim_time, scrim_status FROM scrim INNER JOIN team ON scrim.team_id = team.team_id WHERE scrim.scrim_status = $1 and scrim.scrim_map = $2 AND scrim.game_id = (SELECT game_id FROM team WHERE team_id = $3) AND (scrim.scrim_date > CURRENT_DATE OR (scrim.scrim_date = CURRENT_DATE  AND scrim.scrim_time > CURRENT_TIME )) ORDER BY scrim.scrim_date ASC, scrim.scrim_time ASC;", "unmatched", data.ScrimMap, data.TeamId)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = database.Db.Query("SELECT scrim.scrim_id, scrim.team_id, team.team_logo, team.team_name, scrim.scrim_map, scrim.scrim_date, scrim.scrim_time, scrim_status FROM scrim INNER JOIN team ON scrim.team_id = team.team_id WHERE scrim.scrim_status = $1 AND scrim.game_id = (SELECT game_id FROM team WHERE team_id = $2) AND (scrim.scrim_date > CURRENT_DATE OR (scrim.scrim_date = CURRENT_DATE  AND scrim.scrim_time > CURRENT_TIME )) ORDER BY scrim.scrim_date ASC, scrim.scrim_time ASC;", "unmatched", data.TeamId)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()
	for rows.Next() {
		var detail model.ScrimDetailForQuery
		err := rows.Scan(&detail.ScrimId, &detail.TeamId, &detail.TeamLogo, &detail.TeamName, &detail.ScrimMap, &detail.ScrimDate, &detail.ScrimTime, &detail.ScrimStatus)
		if err != nil {
			return nil, err
		}
		scrims.Scrims = append(scrims.Scrims, detail)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// map flag
	var scrimOffers model.ScrimOffer
	scrimOfferRows, err := database.Db.Query("SELECT scrim_id FROM scrim_offer WHERE team_id = $1 and offer_status = 'pending';", data.TeamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for scrimOfferRows.Next() {
		var detail model.ScrimOfferDetail
		err := scrimOfferRows.Scan(&detail.ScrimId)
		if err != nil {
			return nil, err
		}
		scrimOffers.ScrimOffers = append(scrimOffers.ScrimOffers, detail)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	for i, v := range scrims.Scrims {
		if v.TeamId == data.TeamId {
			scrims.Scrims[i].Flag = "delete"
			continue
		}

		for _, k := range scrimOffers.ScrimOffers {
			fmt.Printf("%v :: %v", v.ScrimId, k.ScrimId)
			if v.ScrimId == k.ScrimId {
				scrims.Scrims[i].Flag = "withdraw offer"
				break
			}
		}

		if (scrims.Scrims[i].Flag != "delete") && (scrims.Scrims[i].Flag != "withdraw offer") {
			scrims.Scrims[i].Flag = "make offer"
		}
	}

	return &scrims, nil
}

func ScrimGetMatch(teamId string) (*model.ScrimGet, error) {
	var scrims model.ScrimGet

	rows, err := database.Db.Query("SELECT scrim.scrim_id, team.team_id, team.team_logo, team.team_name, scrim.scrim_map, scrim.scrim_date, scrim.scrim_time, scrim_status FROM scrim INNER JOIN team ON (scrim.team_id = team.team_id OR scrim.offer_team_id = team.team_id)  WHERE (scrim.team_id = $1 or scrim.offer_team_id = $2) and scrim.scrim_status = 'matched' AND (scrim.scrim_date > CURRENT_DATE OR (scrim.scrim_date = CURRENT_DATE  AND scrim.scrim_time > CURRENT_TIME )) ORDER BY scrim.scrim_date ASC, scrim.scrim_time ASC;", teamId, teamId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var detail model.ScrimDetail
		err := rows.Scan(&detail.ScrimId, &detail.TeamId, &detail.TeamLogo, &detail.TeamName, &detail.ScrimMap, &detail.ScrimDate, &detail.ScrimTime, &detail.ScrimStatus)
		if err != nil {
			return nil, err
		}

		teamId, err := strconv.Atoi(teamId)
		if err != nil {
			return nil, err
		}
		if detail.TeamId == teamId {
			scrims.Scrims = append(scrims.Scrims, detail)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &scrims, nil

}
