package shared

import (
	"context"

	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

type MailSlurpClient struct {
	API            *mailslurp.APIClient
	DefaultInboxID string

	apiKey string
}

func NewMailSlurpClient(apiKey, defaultInboxID string) *MailSlurpClient {
	cfg := mailslurp.NewConfiguration()
	cfg.AddDefaultHeader("x-api-key", apiKey)

	return &MailSlurpClient{
		API:            mailslurp.NewAPIClient(cfg),
		DefaultInboxID: defaultInboxID,
		apiKey:         apiKey,
	}
}

// AuthContext returns a context carrying the MailSlurp API key, required by
// SDK methods that read auth from context (ContextAPIKey).
func (c *MailSlurpClient) AuthContext(parent context.Context) context.Context {
	return context.WithValue(parent, mailslurp.ContextAPIKey, mailslurp.APIKey{Key: c.apiKey})
}
