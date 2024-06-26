package service

import (
	"log"
	"scrim-api/database"
	send_email "scrim-api/email_test"
	"scrim-api/model"
)

func UserRegister(info model.UserRegisterReq) error {
	_, err := database.Db.Exec("INSERT INTO \"user\" (user_pass, nickname, phone_number, email, game_id) VALUES ($1, $2, $3, $4, $5)",
		info.UserPass,
		info.Nickname,
		info.PhoneNumber,
		info.Email,
		info.GameId)
	if err != nil {
		return err
	}

	return nil
}

func UserLogin(data model.UserLoginReq) (*model.UserData, error) {
	var userData model.UserData

	err := database.Db.QueryRow(`
    SELECT 
        u.user_id, 
        u.nickname, 
        u.phone_number, 
        u.email, 
		u.role,
        u.game_id, 
        g.game_name,
        u.team_id,
        t.team_name,
		t.team_logo
    FROM 
        "user" u
    JOIN 
        game g ON u.game_id = g.game_id
    LEFT JOIN 
        team t ON u.team_id = t.team_id
    WHERE 
        u.email = $1 AND u.user_pass = $2`,
		data.Email,
		data.UserPass).Scan(
		&userData.UserId,
		&userData.Nickname,
		&userData.PhoneNumber,
		&userData.Email,
		&userData.Role,
		&userData.GameId,
		&userData.GameName,
		&userData.TeamId,
		&userData.TeamName,
		&userData.TeamLogo,
	)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func ChangeRole(data model.ChangeRole) error {
	_, err := database.Db.Exec("UPDATE \"user\" SET role = $1 WHERE user_id = $2;",
		data.Role, data.UserId)
	if err != nil {
		return err
	}

	return nil
}

func KickMember(data model.KickMember) error {
	_, err := database.Db.Exec("UPDATE \"user\" SET role = $1, team_id = $3 WHERE user_id = $2;",
		"Player", data.UserId, nil)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserProfile(user_id string, data model.UserUpdateData) error {
	var err error
	if data.UserPass != nil {
		_, err = database.Db.Exec("UPDATE \"user\" SET user_pass = $1 WHERE user_id = $2;", data.UserPass, user_id)	
	}
	if err != nil {
		return err
	}
	
	if data.Nickname != nil {
		_, err = database.Db.Exec("UPDATE \"user\" SET nickname = $1 WHERE user_id = $2;", data.Nickname, user_id)	
	}
	if err != nil {
		return err
	}

	if data.PhoneNumber != nil {
		_, err = database.Db.Exec("UPDATE \"user\" SET phone_number = $1 WHERE user_id = $2;", data.PhoneNumber, user_id)	
	}
	if err != nil {
		return err
	}

	return nil
}

func SendPasswordToEmail(data model.ForgotPasswordReq) error {
	var userpass string
	err := database.Db.QueryRow(`
    SELECT 
		user_pass
    FROM 
        "user"
    WHERE 
        email = $1`,
		data.Email).Scan(
		&userpass,
	)
	if err != nil {
		return err
	}

	// send email
	// if err := email_service.ES.SendPasswordtoEmail(data.Email, userpass); err != nil {
	// 	log.Println("Error sending email:", err)
	// 	return err
	// }
	if err := send_email.SendEmail(data.Email, userpass); err != nil {
		log.Println("Error sending email:", err)
		return err
	}
	log.Println("Successfully sent email")

	return nil
}
