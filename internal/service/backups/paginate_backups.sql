-- name: BackupsServicePaginateBackupsCount :one
SELECT COUNT(*) FROM backups;

-- name: BackupsServicePaginateBackups :many
SELECT *
FROM backups
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
