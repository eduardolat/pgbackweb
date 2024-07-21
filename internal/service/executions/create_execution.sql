-- name: ExecutionsServiceCreateExecution :one
INSERT INTO executions (status, message, path)
VALUES (@status, @message, @path)
RETURNING *;
