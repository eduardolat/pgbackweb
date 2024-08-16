-- name: BackupsServiceGetBackupsQty :one
SELECT 
  COUNT(*) AS all,
  SUM(CASE WHEN is_active = true THEN 1 ELSE 0 END) AS active,
  SUM(CASE WHEN is_active = false THEN 1 ELSE 0 END) AS inactive
FROM backups;
