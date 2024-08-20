-- name: BackupsServiceGetAllBackups :many
SELECT * FROM backups
ORDER BY created_at DESC;
