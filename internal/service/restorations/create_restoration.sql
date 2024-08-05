-- name: RestorationsServiceCreateRestoration :one
INSERT INTO restorations (execution_id, database_id, status, message)
VALUES (@execution_id, @database_id, @status, @message)
RETURNING *;
