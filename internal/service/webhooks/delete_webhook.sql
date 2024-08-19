-- name: WebhooksServiceDeleteWebhook :exec
DELETE FROM webhooks WHERE id = @webhook_id;
