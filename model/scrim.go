package model

type ScrimPost struct {
	ScrimDate string `json:"scrim_date" binding:"required"`
	ScrimTime string `json:"scrim_time" binding:"required"`
	ScrimMap  string `json:"scrim_map" binding:"required"`
	TeamId    int    `json:"team_id" binding:"required"`
	GameId    int    `json:"game_id" binding:"required"`
}

type ScrimMakeOffer struct {
	ScrimId int `json:"scrim_id" binding:"required"`
	TeamId  int `json:"team_id" binding:"required"`
}

type ScrimAcceptOffer struct {
	ScrimId int `json:"scrim_id" binding:"required"`
	TeamId  int `json:"team_id" binding:"required"`
}

type ScrimCancelMatch struct {
	ScrimId int `json:"scrim_id" binding:"required"`
	TeamId  int `json:"team_id" binding:"required"`
}

type ScrimDelete struct {
	ScrimId int `json:"scrim_id" binding:"required"`
}

type ScrimGetReq struct {
	ScrimMap *string `json:"scrim_map"`
}

type ScrimGetAll struct {
	Scrims []ScrimDetail `json:"scrims"`
}

type ScrimDetail struct {
	ScrimId   int    `json:"scrim_id"`
	TeamLogo  string `json:"team_logo"`
	TeamName  string `json:"team_name"`
	ScrimMap  string `json:"scrim_map"`
	ScrimDate string `json:"scrim_date"`
	ScrimTime string `json:"scrim_time"`
}
