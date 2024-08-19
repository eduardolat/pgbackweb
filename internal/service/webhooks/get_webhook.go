package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) GetWebhook(
	ctx context.Context, id uuid.UUID,
) (dbgen.Webhook, error) {
	return s.dbgen.WebhooksServiceGetWebhook(ctx, id)
}
