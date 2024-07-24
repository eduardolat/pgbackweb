-- name: ExecutionsServicePaginateExecutionsCount :one
SELECT COUNT(executions.*)
FROM executions
JOIN backups ON backups.id = executions.backup_id
JOIN databases ON databases.id = backups.database_id
JOIN destinations ON destinations.id = backups.destination_id
WHERE
(
  sqlc.narg('backup_id')::UUID IS NULL
  OR
  backups.id = sqlc.narg('backup_id')::UUID
)
AND
(
  sqlc.narg('database_id')::UUID IS NULL
  OR
  databases.id = sqlc.narg('database_id')::UUID
)
AND
(
  sqlc.narg('destination_id')::UUID IS NULL
  OR
  destinations.id = sqlc.narg('destination_id')::UUID
);

-- name: ExecutionsServicePaginateExecutions :many
SELECT
  executions.*,
  backups.name AS backup_name,
  databases.name AS database_name,
  destinations.name AS destination_name
FROM executions
JOIN backups ON backups.id = executions.backup_id
JOIN databases ON databases.id = backups.database_id
JOIN destinations ON destinations.id = backups.destination_id
WHERE
(
  sqlc.narg('backup_id')::UUID IS NULL
  OR
  backups.id = sqlc.narg('backup_id')::UUID
)
AND
(
  sqlc.narg('database_id')::UUID IS NULL
  OR
  databases.id = sqlc.narg('database_id')::UUID
)
AND
(
  sqlc.narg('destination_id')::UUID IS NULL
  OR
  destinations.id = sqlc.narg('destination_id')::UUID
)
ORDER BY executions.started_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
