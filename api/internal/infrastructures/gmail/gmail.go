package gmail

import (
	"github.com/boke0ya/beathub-api/internal/adapters"
	"github.com/boke0ya/beathub-api/internal/errors"
	"gopkg.in/gomail.v2"
)

type GMailAdapter struct {
	email    string
	password string
}

func NewGMailAdapter(email string, password string) adapters.MailAdapter {
	return GMailAdapter{
		email:    email,
		password: password,
	}
}

func (mailer GMailAdapter) Send(to string, subject string, body string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", mailer.email)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 465, mailer.email, mailer.password)
	if err := d.DialAndSend(mail); err != nil {
		return errors.New(errors.FailedToSendEmailError, err)
	} else {
		return nil
	}
}
