package webhooks

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// RunDatabaseHealthy runs the healthy webhooks for the given database ID.
func (s *Service) RunDatabaseHealthy(databaseID uuid.UUID, databaseStatus DatabaseStatus) {
	go func() {
		ctx := context.Background()
		webhookPayload := WebhookPayload{
			Database: databaseStatus,
		}
		runWebhook(s, ctx, EventTypeDatabaseHealthy, databaseID, webhookPayload)
	}()
}

// RunDatabaseUnhealthy runs the unhealthy webhooks for the given database ID.
func (s *Service) RunDatabaseUnhealthy(databaseID uuid.UUID, databaseStatus DatabaseStatus) {
	go func() {
		webhookPayload := WebhookPayload{
			Database: databaseStatus,
		}
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDatabaseUnhealthy, databaseID, webhookPayload)
	}()
}

// RunDestinationHealthy runs the healthy webhooks for the given destination ID.
func (s *Service) RunDestinationHealthy(destinationID uuid.UUID, destinationStatus DestinationStatus) {
	go func() {
		webhookPayload := WebhookPayload{
			Destination: destinationStatus,
		}
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDestinationHealthy, destinationID, webhookPayload)
	}()
}

// RunDestinationUnhealthy runs the unhealthy webhooks for the given
// destination ID.
func (s *Service) RunDestinationUnhealthy(destinationID uuid.UUID, destinationStatus DestinationStatus) {
	go func() {
		webhookPayload := WebhookPayload{
			Destination: destinationStatus,
		}

		ctx := context.Background()
		runWebhook(s, ctx, EventTypeDestinationUnhealthy, destinationID, webhookPayload)
	}()
}

// RunExecutionSuccess runs the success webhooks for the given execution ID.
func (s *Service) RunExecutionSuccess(backupID uuid.UUID, webhookPayload WebhookPayload) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeExecutionSuccess, backupID, webhookPayload)
	}()
}

// RunExecutionFailed runs the failed webhooks for the given execution ID.
func (s *Service) RunExecutionFailed(backupID uuid.UUID, webhookPayload WebhookPayload) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, EventTypeExecutionFailed, backupID, webhookPayload)
	}()
}

// runWebhook runs the webhooks for the given event type and target ID.
func runWebhook(
	s *Service, ctx context.Context, eventType eventType, targetID uuid.UUID,
	webhookPayload WebhookPayload,
) {

	webhooks, err := s.dbgen.WebhooksServiceGetWebhooksToRun(
		ctx, dbgen.WebhooksServiceGetWebhooksToRunParams{
			EventType: eventType.Value.Key,
			TargetID:  targetID,
		},
	)
	webhookPayload.EventType = eventType.Value.Name
	webhookPayload.Msg = buildMessage(eventType, webhookPayload)
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

			err := s.SendWebhookRequest(ctx, webhook, webhookPayload)
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
