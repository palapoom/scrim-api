package service

import (
	"scrim-api/database"
	"scrim-api/model"
)

func MapNameGet(gameId string) (*model.MapGet, error) {
	var mapsName []string
	rows, err := database.Db.Query("SELECT map_name FROM map WHERE game_id = $1", gameId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		mapsName = append(mapsName, name)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &model.MapGet{
		Maps: mapsName,
	}, nil
}
