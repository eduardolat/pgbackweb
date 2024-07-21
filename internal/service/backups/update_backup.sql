-- name: BackupsServiceUpdateBackup :one
UPDATE backups
SET
  database_id = @database_id,
  destination_id = @destination_id,
  name = @name,
  cron_expression = @cron_expression,
  is_active = @is_active,
  dest_dir = @dest_dir,
  opt_data_only = @opt_data_only,
  opt_schema_only = @opt_schema_only,
  opt_clean = @opt_clean,
  opt_if_exists = @opt_if_exists,
  opt_create = @opt_create,
  opt_no_comments = @opt_no_comments
WHERE id = @id
RETURNING *;
