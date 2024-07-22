-- name: ExecutionsServiceGetExpiredExecutions :many
SELECT executions.*
FROM executions
JOIN backups ON executions.backup_id = backups.id
WHERE
  backups.retention_days > 0
  AND executions.status != 'deleted'
  AND executions.finished_at IS NOT NULL
  AND (
    executions.finished_at + (backups.retention_days || ' days')::INTERVAL
  ) < NOW();
