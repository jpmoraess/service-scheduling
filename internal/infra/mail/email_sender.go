package mail

import (
	"github.com/jpmoraess/service-scheduling/internal/domain"
	"gopkg.in/gomail.v2"
)

type GoEmailSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewGoEmailSender(host string, port int, username, password string) *GoEmailSender {
	return &GoEmailSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (s *GoEmailSender) Send(email domain.Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)

	return d.DialAndSend(m)
}
