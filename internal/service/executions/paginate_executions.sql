-- name: ExecutionsServicePaginateExecutionsCount :one
SELECT COUNT(executions.*)
FROM executions
INNER JOIN backups ON backups.id = executions.backup_id
INNER JOIN databases ON databases.id = backups.database_id
LEFT JOIN destinations ON destinations.id = backups.destination_id
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
  databases.pg_version AS database_pg_version,
  destinations.name AS destination_name,
  backups.is_local AS backup_is_local
FROM executions
INNER JOIN backups ON backups.id = executions.backup_id
INNER JOIN databases ON databases.id = backups.database_id
LEFT JOIN destinations ON destinations.id = backups.destination_id
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
