package delivery

import "time"

type SendEmailResult struct {
	ID      string    `json:"id"`
	InboxID string    `json:"inboxId"`
	To      []string  `json:"to"`
	Subject string    `json:"subject"`
	SentAt  time.Time `json:"sentAt"`
}
