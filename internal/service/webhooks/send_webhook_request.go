package webhooks

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
)

// SendWebhookRequest sends a webhook request to the given webhook and
// stores the result in the database.
func (s *Service) SendWebhookRequest(
	ctx context.Context, webhook dbgen.Webhook,
) error {
	timeStart := time.Now()

	if !webhook.Body.Valid || webhook.Body.String == "" {
		webhook.Body = sql.NullString{String: "{}", Valid: true}
	}
	bodyReader := strings.NewReader(webhook.Body.String)

	if !webhook.Headers.Valid || webhook.Headers.String == "" {
		webhook.Headers = sql.NullString{String: "{}", Valid: true}
	}
	headers := map[string]string{}
	err := json.Unmarshal([]byte(webhook.Headers.String), &headers)
	if err != nil {
		return fmt.Errorf("error parsing headers: %w", err)
	}

	client := http.Client{Timeout: time.Second * 30}
	req, err := http.NewRequestWithContext(
		ctx, webhook.Method, webhook.Url, bodyReader,
	)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	resHeaders, err := json.Marshal(res.Header)
	if err != nil {
		return fmt.Errorf("error marshalling response headers: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	_, err = s.dbgen.WebhooksServiceCreateWebhookExecution(
		ctx, dbgen.WebhooksServiceCreateWebhookExecutionParams{
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
		return fmt.Errorf("error updating webhook result: %w", err)
	}

	logger.Info("webhook sent successfully", logger.KV{
		"webhook_id": webhook.ID,
		"status":     res.Status,
	})

	return nil
}
