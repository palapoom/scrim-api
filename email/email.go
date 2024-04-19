package email_service

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

// SendPasswordResetEmail sends a password reset email.
func (es *EmailService) SendPasswordtoEmail(to, code string) error {
	t, err := template.New("email").Parse("<h2>Your WeScrim Account Password</h2><br><p>This is your password for WeScrim account is <b> {{.Password}} </b></p><p>Please change your password as soon as possible for security reasons.</p><br><p>Thanks, WeScrim team</p>")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your WeScrim Account Password | WeScrim \n%s\n\n", mimeHeaders)))

	if err := t.Execute(&body, struct{ Password string }{Password: code}); err != nil {
		return err
	}

	return smtp.SendMail(fmt.Sprintf("%s:%s", es.SMTPHost, es.SMTPPort), Auth, es.Username, []string{to}, body.Bytes())
}
