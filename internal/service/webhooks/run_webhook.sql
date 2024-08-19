-- name: WebhooksServiceGetWebhooksToRun :many
SELECT * FROM webhooks
WHERE is_active = true
AND event_type = @event_type
AND @target_id = ANY(target_ids);

-- name: WebhooksServiceCreateWebhookResult :one
INSERT INTO webhook_results (
  webhook_id, req_method, req_headers, req_body,
  res_status, res_headers, res_body, res_duration
)
VALUES (
  @webhook_id, @req_method, @req_headers, @req_body,
  @res_status, @res_headers, @res_body, @res_duration
)
RETURNING *;
