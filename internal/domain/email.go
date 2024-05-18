package domain

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

type EmailSender interface {
	Send(email Email) error
}
