-- name: WebhooksServiceGetWebhook :one
SELECT * FROM webhooks WHERE id = @webhook_id;
