package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateWebhook(
	ctx context.Context, params dbgen.WebhooksServiceCreateWebhookParams,
) (dbgen.Webhook, error) {
	return s.dbgen.WebhooksServiceCreateWebhook(ctx, params)
}
