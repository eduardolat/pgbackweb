-- name: ExecutionsServiceGetExecutionsQty :one
SELECT 
  COUNT(*) AS all,
  COALESCE(SUM(CASE WHEN status = 'running' THEN 1 ELSE 0 END), 0)::INTEGER AS running,
  COALESCE(SUM(CASE WHEN status = 'success' THEN 1 ELSE 0 END), 0)::INTEGER AS success,
  COALESCE(SUM(CASE WHEN status = 'failed' THEN 1 ELSE 0 END), 0)::INTEGER AS failed,
  COALESCE(SUM(CASE WHEN status = 'deleted' THEN 1 ELSE 0 END), 0)::INTEGER AS deleted
FROM executions;
