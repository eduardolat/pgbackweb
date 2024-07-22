-- name: BackupsServiceGetScheduleAllData :many
SELECT 
  id,
  is_active,
  cron_expression,
  time_zone
FROM backups
ORDER BY created_at DESC;
