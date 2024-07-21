-- name: ExecutionsServiceCreateExecution :one
INSERT INTO executions (backup_id, status, message, path)
VALUES (@backup_id, @status, @message, @path)
RETURNING *;
