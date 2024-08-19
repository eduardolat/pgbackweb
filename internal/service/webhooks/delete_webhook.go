package webhooks

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteWebhook(
	ctx context.Context, id uuid.UUID,
) error {
	return s.dbgen.WebhooksServiceDeleteWebhook(ctx, id)
}
