-- name: ExecutionsServiceGetExecution :one
SELECT * FROM executions
WHERE id = @id;
