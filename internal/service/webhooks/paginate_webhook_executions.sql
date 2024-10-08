-- name: WebhooksServicePaginateWebhookExecutionsCount :one
SELECT COUNT(*) FROM webhook_executions
WHERE webhook_id = @webhook_id;

-- name: WebhooksServicePaginateWebhookExecutions :many
SELECT * FROM webhook_executions
WHERE webhook_id = @webhook_id
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');