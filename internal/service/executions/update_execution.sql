-- name: ExecutionsServiceUpdateExecution :one
UPDATE executions
SET
  status = @status,
  message = @message,
  path = @path,
  finished_at = @finished_at,
  deleted_at = @deleted_at
WHERE id = @id
RETURNING *;
