-- name: WebhooksServiceCreateWebhook :one
INSERT INTO webhooks (
  name, is_active, event_type, target_ids,
  url, method, headers, body
) VALUES (
  @name, @is_active, @event_type, @target_ids,
  @url, @method, @headers, @body
) RETURNING *;
