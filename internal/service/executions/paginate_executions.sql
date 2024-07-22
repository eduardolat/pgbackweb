-- name: ExecutionsServicePaginateExecutionsCount :one
SELECT COUNT(*) FROM executions;

-- name: ExecutionsServicePaginateExecutions :many
SELECT *
FROM executions
ORDER BY started_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
