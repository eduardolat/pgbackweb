-- name: WebhooksServicePaginateWebhooksCount :one
SELECT COUNT(*) FROM webhooks;

-- name: WebhooksServicePaginateWebhooks :many
SELECT * FROM webhooks
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
