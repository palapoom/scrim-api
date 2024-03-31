package service

import (
	"fmt"
	"scrim-api/database"
	"scrim-api/model"
)

func ScrimPost(data model.ScrimPost) (*int, error) {
	var lastScrimId int
	err := database.Db.QueryRow("SELECT scrim_id FROM scrim ORDER BY scrim_id DESC LIMIT 1").Scan(&lastScrimId)
	if err != nil {
		return nil, err
	}

	fmt.Printf("last scrim_id: %v", lastScrimId)

	_, err = database.Db.Exec("INSERT INTO \"scrim\" (scrim_id, scrim_date, scrim_time, scrim_map, game_id, team_id) VALUES ($1, $2, $3, $4, $5, $6)",
		lastScrimId+1, data.ScrimDate, data.ScrimTime, data.ScrimMap, data.GameId, data.TeamId)

	if err != nil {
		return nil, err
	}

	lastScrimId = lastScrimId + 1

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
