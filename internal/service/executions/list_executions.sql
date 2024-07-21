-- name: ExecutionsServiceListExecutions :many
SELECT * FROM executions ORDER BY started_at DESC;
