package adapters

type MailAdapter interface {
	Send(to string, subject string, content string) error
}
