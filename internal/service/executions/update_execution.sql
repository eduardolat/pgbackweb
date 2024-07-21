-- name: ExecutionsServiceUpdateExecution :one
UPDATE executions
SET
  status = COALESCE(status, @status),
  message = COALESCE(message, @message),
  path = COALESCE(path, @path),
  finished_at = COALESCE(finished_at, @finished_at),
  deleted_at = COALESCE(deleted_at, @deleted_at)
WHERE id = @id
RETURNING *;
