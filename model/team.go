package model

type TeamCreateReq struct {
	TeamName string  `json:"team_name" binding:"required"`
	TeamLogo *string `json:"team_logo"`
	GameId   int     `json:"game_id" binding:"required"`
}

type TeamUpdate struct {
	TeamId   int     `json:"team_id" binding:"required"`
	TeamName string  `json:"team_name" binding:"required"`
	TeamLogo *string `json:"team_logo"`
}

type TeamJoin struct {
	InviteCode string `json:"invite_code" binding:"required"`
	UserId     string `json:"user_id" binding:"required"`
}

type TeamMember struct {
	Members []Member `json:"members"`
}

type Member struct {
	UserId   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
}

type TeamDetail struct {
	TeamName   string   `json:"team_name"`
	GameId     int      `json:"game_id"`
	TeamLogo   string   `json:"team_logo"`
	InviteCode string   `json:"invite_code"`
	InviteFlag bool     `json:"invite_flag"`
	Members    []Member `json:"members"`
}

type TeamInviteCodeGet struct { // set flag response
	InviteCode string `json:"invite_code"`
}
