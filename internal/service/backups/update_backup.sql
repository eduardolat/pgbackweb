-- name: BackupsServiceUpdateBackup :one
UPDATE backups
SET
  name = COALESCE(name, @name),
  cron_expression = COALESCE(cron_expression, @cron_expression),
  time_zone = COALESCE(time_zone, @time_zone),
  is_active = COALESCE(is_active, @is_active),
  dest_dir = COALESCE(dest_dir, @dest_dir),
  opt_data_only = COALESCE(opt_data_only, @opt_data_only),
  opt_schema_only = COALESCE(opt_schema_only, @opt_schema_only),
  opt_clean = COALESCE(opt_clean, @opt_clean),
  opt_if_exists = COALESCE(opt_if_exists, @opt_if_exists),
  opt_create = COALESCE(opt_create, @opt_create),
  opt_no_comments = COALESCE(opt_no_comments, @opt_no_comments)
WHERE id = @id
RETURNING *;
