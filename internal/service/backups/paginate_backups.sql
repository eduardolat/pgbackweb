-- name: BackupsServicePaginateBackupsCount :one
SELECT COUNT(*) FROM backups;

-- name: BackupsServicePaginateBackups :many
SELECT
  backups.*,
  databases.name AS database_name,
  destinations.name AS destination_name
FROM backups
JOIN databases ON backups.database_id = databases.id
JOIN destinations ON backups.destination_id = destinations.id
ORDER BY backups.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
