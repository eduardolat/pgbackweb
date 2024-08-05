-- name: RestorationsServicePaginateRestorationsCount :one
SELECT COUNT(restorations.*)
FROM restorations
INNER JOIN executions ON executions.id = restorations.execution_id
LEFT JOIN databases ON databases.id = restorations.database_id
WHERE
(
  sqlc.narg('execution_id')::UUID IS NULL
  OR
  restorations.execution_id = sqlc.narg('execution_id')::UUID
)
AND
(
  sqlc.narg('database_id')::UUID IS NULL
  OR
  restorations.database_id = sqlc.narg('database_id')::UUID
);

-- name: RestorationsServicePaginateRestorations :many
SELECT
  restorations.*,
  databases.name AS database_name
FROM restorations
INNER JOIN executions ON executions.id = restorations.execution_id
LEFT JOIN databases ON databases.id = restorations.database_id
WHERE
(
  sqlc.narg('execution_id')::UUID IS NULL
  OR
  restorations.execution_id = sqlc.narg('execution_id')::UUID
)
AND
(
  sqlc.narg('database_id')::UUID IS NULL
  OR
  restorations.database_id = sqlc.narg('database_id')::UUID
)
ORDER BY restorations.started_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
