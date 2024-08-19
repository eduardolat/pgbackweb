-- name: WebhooksServiceUpdateWebhook :one
UPDATE webhooks
SET
  name = COALESCE(sqlc.narg('name'), name),
  is_active = COALESCE(sqlc.narg('is_active'), is_active),
  target_ids = COALESCE(sqlc.narg('target_ids'), target_ids),
  url = COALESCE(sqlc.narg('url'), url),
  method = COALESCE(sqlc.narg('method'), method),
  headers = COALESCE(sqlc.narg('headers'), headers),
  body = COALESCE(sqlc.narg('body'), body)
WHERE id = @webhook_id
RETURNING *;