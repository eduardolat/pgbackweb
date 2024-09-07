-- name: BackupsServiceGetBackupsQty :one
SELECT 
  COUNT(*) AS all,
  COALESCE(SUM(CASE WHEN is_active = true THEN 1 ELSE 0 END), 0)::INTEGER AS active,
  COALESCE(SUM(CASE WHEN is_active = false THEN 1 ELSE 0 END), 0)::INTEGER AS inactive
FROM backups;
