package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) DuplicateWebhook(
	ctx context.Context, webhookID uuid.UUID,
) (dbgen.Webhook, error) {
	return s.dbgen.WebhooksServiceDuplicateWebhook(ctx, webhookID)
}
