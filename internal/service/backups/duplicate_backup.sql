-- name: BackupsServiceDuplicateBackup :one
INSERT INTO backups
SELECT (
  backups
  #= hstore('id', uuid_generate_v4()::text)
  #= hstore('name', (backups.name || ' (copy)')::text)
  #= hstore('is_active', false::text)
  #= hstore('created_at', now()::text)
  #= hstore('updated_at', now()::text)
).*
FROM backups
WHERE backups.id = @backup_id
RETURNING *;
