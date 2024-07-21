-- name: BackupsServiceGetScheduleAllData :many
SELECT 
  id,
  is_active,
  cron_expression,
  time_zone
FROM backups
WHERE is_active = TRUE
ORDER BY created_at DESC;
