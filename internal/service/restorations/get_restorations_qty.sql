-- name: RestorationsServiceGetRestorationsQty :one
SELECT 
  COUNT(*) AS all,
  SUM(CASE WHEN status = 'running' THEN 1 ELSE 0 END) AS running,
  SUM(CASE WHEN status = 'success' THEN 1 ELSE 0 END) AS success,
  SUM(CASE WHEN status = 'failed' THEN 1 ELSE 0 END) AS failed
FROM restorations;
