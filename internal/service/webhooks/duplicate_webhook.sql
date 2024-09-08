-- name: WebhooksServiceDuplicateWebhook :one
INSERT INTO webhooks
SELECT (
  webhooks
  #= hstore('id', uuid_generate_v4()::text)
  #= hstore('name', (webhooks.name || ' (copy)')::text)
  #= hstore('is_active', false::text)
  #= hstore('created_at', now()::text)
  #= hstore('updated_at', now()::text)
).*
FROM webhooks
WHERE webhooks.id = @webhook_id
RETURNING *;
