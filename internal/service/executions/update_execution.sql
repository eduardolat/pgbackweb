-- name: ExecutionsServiceUpdateExecution :one
UPDATE executions
SET
  status = COALESCE(sqlc.narg('status'), status),
  message = COALESCE(sqlc.narg('message'), message),
  path = COALESCE(sqlc.narg('path'), path),
  finished_at = COALESCE(sqlc.narg('finished_at'), finished_at),
  deleted_at = COALESCE(sqlc.narg('deleted_at'), deleted_at)
WHERE id = @id
RETURNING *;
