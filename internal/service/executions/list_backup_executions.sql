-- name: ExecutionsServiceListBackupExecutions :many
SELECT * FROM executions
WHERE backup_id = @backup_id
ORDER BY started_at DESC;
