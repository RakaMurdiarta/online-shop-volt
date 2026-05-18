package delivery

type SendEmailRequest struct {
	InboxID  string
	To       []string
	Cc       []string
	Bcc      []string
	From     string
	FromName string
	ReplyTo  string
	Subject  string
	Body     string
	IsHTML   bool
}
