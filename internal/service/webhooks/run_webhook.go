package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// RunDatabaseHealthy runs the healthy webhooks for the given database ID.
func (s *Service) RunDatabaseHealthy(databaseID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDatabaseHealthy, databaseID)
	}()
}

// RunDatabaseUnhealthy runs the unhealthy webhooks for the given database ID.
func (s *Service) RunDatabaseUnhealthy(databaseID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDatabaseUnhealthy, databaseID)
	}()
}

// RunDestinationHealthy runs the healthy webhooks for the given destination ID.
func (s *Service) RunDestinationHealthy(destinationID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDestinationHealthy, destinationID)
	}()
}

// RunDestinationUnhealthy runs the unhealthy webhooks for the given
// destination ID.
func (s *Service) RunDestinationUnhealthy(destinationID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDestinationUnhealthy, destinationID)
	}()
}

// RunExecutionSuccess runs the success webhooks for the given execution ID.
func (s *Service) RunExecutionSuccess(backupID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeExecutionSuccess, backupID)
	}()
}

// RunExecutionFailed runs the failed webhooks for the given execution ID.
func (s *Service) RunExecutionFailed(backupID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeExecutionFailed, backupID)
	}()
}

// runWebhook runs the webhooks for the given event type and target ID.
func runWebhook(
	s *Service, ctx context.Context, eventType eventType, targetID uuid.UUID,
) {
	webhooks, err := s.dbgen.WebhooksServiceGetWebhooksToRun(
		ctx, dbgen.WebhooksServiceGetWebhooksToRunParams{
			EventType: eventType.Value.Key,
			TargetID:  targetID,
		},
	)
	if err != nil {
		logger.Error("error getting webhooks to run", logger.KV{"error": err})
		return
	}
	if len(webhooks) == 0 {
		return
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(5)

	for _, webhook := range webhooks {
		eg.Go(func() error {
			err := s.SendWebhookRequest(ctx, webhook)
			if err != nil {
				logger.Error("error sending webhook request", logger.KV{
					"webhook_id": webhook.ID,
					"error":      err.Error(),
				})
			}
			return nil
		})
	}

	_ = eg.Wait()
}
