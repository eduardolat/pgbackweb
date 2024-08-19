package webhooks

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// RunDatabaseHealthy runs the healthy webhooks for the given database ID.
func (s *Service) RunDatabaseHealthy(databaseID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, eventTypeDatabaseHealthy, databaseID)
	}()
}

// RunDatabaseUnhealthy runs the unhealthy webhooks for the given database ID.
func (s *Service) RunDatabaseUnhealthy(databaseID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, eventTypeDatabaseUnhealthy, databaseID)
	}()
}

// RunDestinationHealthy runs the healthy webhooks for the given destination ID.
func (s *Service) RunDestinationHealthy(destinationID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, eventTypeDestinationHealthy, destinationID)
	}()
}

// RunDestinationUnhealthy runs the unhealthy webhooks for the given
// destination ID.
func (s *Service) RunDestinationUnhealthy(destinationID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, eventTypeDestinationUnhealthy, destinationID)
	}()
}

// RunExecutionSuccess runs the success webhooks for the given execution ID.
func (s *Service) RunExecutionSuccess(backupID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, eventTypeExecutionSuccess, backupID)
	}()
}

// RunExecutionFailed runs the failed webhooks for the given execution ID.
func (s *Service) RunExecutionFailed(backupID uuid.UUID) {
	go func() {
		ctx := context.Background()
		runWebhook(s, ctx, eventTypeExecutionFailed, backupID)
	}()
}

// runWebhook runs the webhooks for the given event type and target ID.
func runWebhook(
	s *Service, ctx context.Context, eventType webhook, targetID uuid.UUID,
) {
	webhooks, err := s.dbgen.WebhooksServiceGetWebhooksToRun(
		ctx, dbgen.WebhooksServiceGetWebhooksToRunParams{
			EventType: eventType.Value,
			TargetID:  []uuid.UUID{targetID},
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
			err := sendWebhookRequest(s, ctx, webhook)
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

// sendWebhookRequest sends a webhook request to the given webhook and
// stores the result in the database.
func sendWebhookRequest(
	s *Service, ctx context.Context, webhook dbgen.Webhook,
) error {
	timeStart := time.Now()

	body := strings.NewReader(webhook.Body.String)
	headers := map[string]string{}
	if webhook.Headers.Valid {
		err := json.Unmarshal([]byte(webhook.Headers.String), &headers)
		if err != nil {
			return err
		}
	}

	client := http.Client{Timeout: time.Second * 30}
	req, err := http.NewRequestWithContext(
		ctx, webhook.Method, webhook.Url, body,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resHeaders, err := json.Marshal(res.Header)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	_, err = s.dbgen.WebhooksServiceCreateWebhookResult(
		ctx, dbgen.WebhooksServiceCreateWebhookResultParams{
			WebhookID:  webhook.ID,
			ReqMethod:  sql.NullString{String: req.Method, Valid: true},
			ReqHeaders: webhook.Headers,
			ReqBody:    webhook.Body,
			ResStatus:  sql.NullInt16{Int16: int16(res.StatusCode), Valid: true},
			ResHeaders: sql.NullString{String: string(resHeaders), Valid: true},
			ResBody:    sql.NullString{String: string(resBody), Valid: true},
			ResDuration: sql.NullInt32{
				Int32: int32(time.Since(timeStart).Milliseconds()),
				Valid: true,
			},
		},
	)
	if err != nil {
		return err
	}

	logger.Info("webhook sent successfully", logger.KV{
		"webhook_id": webhook.ID,
		"status":     res.Status,
	})

	return nil
}
