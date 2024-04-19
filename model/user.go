package model

type UserRegisterReq struct {
	UserPass    string `json:"user_pass" binding:"required"`
	Nickname    string `json:"nickname" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required"`
	GameId      int    `json:"game_id" binding:"required"`
}

type UserLoginReq struct {
	Email    string `json:"email" binding:"required"`
	UserPass string `json:"user_pass" binding:"required"`
}

type UserData struct {
	UserId      string  `json:"user_name"`
	Nickname    string  `json:"nickname"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Role        string  `json:"role"`
	GameId      int     `json:"game_id"`
	GameName    string  `json:"game_name"`
	TeamId      *int    `json:"team_id"`
	TeamName    *string `json:"team_name"`
	TeamLogo    *string `json:"team_logo"`
}

type ChangeRole struct { // by Manager
	UserId string `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}

type KickMember struct {
	UserId string `json:"user_id" binding:"required"`
}

type UserUpdateData struct {
	Nickname    *string  `json:"nickname"`
	PhoneNumber *string `json:"phone_number"`
	UserPass    *string `json:"user_pass"`
}

type ForgotPasswordReq struct {
	Email string `json:"email" binding:"required"`
}

