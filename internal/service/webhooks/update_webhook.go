package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateWebhook(
	ctx context.Context, params dbgen.WebhooksServiceUpdateWebhookParams,
) (dbgen.Webhook, error) {
	return s.dbgen.WebhooksServiceUpdateWebhook(ctx, params)
}
