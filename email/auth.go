package email_service

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

// //// Set Auth Email

var Auth smtp.Auth

func SetAuth(username, password string) {
	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		println(err)
	}
	
	c, err := smtp.NewClient(conn, ES.SMTPHost)
	if err != nil {
		println(err)
	}

	tlsconfig := &tls.Config{
		ServerName: ES.SMTPHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
	}

	auth := LoginAuth(username, password)
	Auth = auth

	if err = c.Auth(Auth); err != nil {
		println(err)
	}
}

////// Set ES

var ES *EmailService

func SetES(es *EmailService) {
	ES = es
	fmt.Println("Successfully created connection to email service")
}

// EmailService represents an email sending service.
type EmailService struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
}

// NewEmailService creates a new instance of EmailService.
func NewEmailService(smtpHost, smtpPort, username, password string) *EmailService {
	return &EmailService{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
	}
}
