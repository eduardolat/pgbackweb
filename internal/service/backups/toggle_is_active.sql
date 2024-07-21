-- name: BackupsServiceToggleIsActive :one
UPDATE backups
SET is_active = NOT is_active
WHERE id = @backup_id
RETURNING *;
