-- name: BackupsServiceUpdateBackup :one
UPDATE backups
SET
  name = COALESCE(name, @name),
  cron_expression = COALESCE(sqlc.narg('cron_expression'), cron_expression),
  time_zone = COALESCE(sqlc.narg('time_zone'), time_zone),
  is_active = COALESCE(sqlc.narg('is_active'), is_active),
  dest_dir = COALESCE(sqlc.narg('dest_dir'), dest_dir),
  retention_days = COALESCE(sqlc.narg('retention_days'), retention_days),
  opt_data_only = COALESCE(sqlc.narg('opt_data_only'), opt_data_only),
  opt_schema_only = COALESCE(sqlc.narg('opt_schema_only'), opt_schema_only),
  opt_clean = COALESCE(sqlc.narg('opt_clean'), opt_clean),
  opt_if_exists = COALESCE(sqlc.narg('opt_if_exists'), opt_if_exists),
  opt_create = COALESCE(sqlc.narg('opt_create'), opt_create),
  opt_no_comments = COALESCE(sqlc.narg('opt_no_comments'), opt_no_comments)
WHERE id = @id
RETURNING *;
