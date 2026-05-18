package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/RakaMurdiarta/online-shop-system/internal/modules/mailer/delivery"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/mailer/services"
	"github.com/RakaMurdiarta/online-shop-system/pkg/shared"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

type mailerServiceImpl struct {
	client *shared.MailSlurpClient
}

func NewMailerService(client *shared.MailSlurpClient) services.MailerService {
	return &mailerServiceImpl{client: client}
}

func (s *mailerServiceImpl) SendEmail(ctx context.Context, req delivery.SendEmailRequest) (*delivery.SendEmailResult, error) {
	inboxID := req.InboxID
	if inboxID == "" {
		inboxID = s.client.DefaultInboxID
	}
	if inboxID == "" {
		return nil, errors.New("mailer: inbox id is required (set req.InboxID or MAILER_INBOX_ID)")
	}
	if len(req.To) == 0 {
		return nil, errors.New("mailer: at least one recipient is required")
	}

	opts := mailslurp.SendEmailOptions{
		To:      &req.To,
		Subject: &req.Subject,
		Body:    &req.Body,
		IsHTML:  &req.IsHTML,
	}
	if len(req.Cc) > 0 {
		opts.Cc = &req.Cc
	}
	if len(req.Bcc) > 0 {
		opts.Bcc = &req.Bcc
	}
	if req.From != "" {
		opts.From = &req.From
	}
	if req.FromName != "" {
		opts.FromName = &req.FromName
	}
	if req.ReplyTo != "" {
		opts.ReplyTo = &req.ReplyTo
	}

	sent, _, err := s.client.API.InboxControllerApi.SendEmailAndConfirm(s.client.AuthContext(ctx), inboxID, opts)
	if err != nil {
		return nil, fmt.Errorf("mailer: send failed: %w", err)
	}

	result := &delivery.SendEmailResult{
		ID:      sent.Id,
		InboxID: sent.InboxId,
		SentAt:  sent.SentAt,
	}
	if sent.To != nil {
		result.To = *sent.To
	}
	if sent.Subject != nil {
		result.Subject = *sent.Subject
	}
	return result, nil
}
