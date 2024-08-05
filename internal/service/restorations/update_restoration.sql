-- name: RestorationsServiceUpdateRestoration :one
UPDATE restorations
SET
  status = COALESCE(sqlc.narg('status'), status),
  message = COALESCE(sqlc.narg('message'), message),
  finished_at = COALESCE(sqlc.narg('finished_at'), finished_at)
WHERE id = @id
RETURNING *;
