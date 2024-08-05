-- name: ExecutionsServiceGetExecution :one
SELECT
  executions.*,
  databases.pg_version AS database_pg_version
FROM executions
INNER JOIN backups ON backups.id = executions.backup_id
INNER JOIN databases ON databases.id = backups.database_id
WHERE id = @id;
