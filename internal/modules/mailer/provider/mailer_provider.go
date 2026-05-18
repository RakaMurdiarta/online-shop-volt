package provider

import (
	mailerServiceImpl "github.com/RakaMurdiarta/online-shop-system/internal/modules/mailer/services/Impl"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/mailer/services"
	"github.com/RakaMurdiarta/online-shop-system/pkg/shared"
)

// MailerProvider constructs the MailerService and returns it so other
// modules (auth verification, order receipts, etc.) can depend on it.
// No HTTP routes are registered here — the mailer is an internal service.
func MailerProvider(mailClient *shared.MailSlurpClient) services.MailerService {
	return mailerServiceImpl.NewMailerService(mailClient)
}
