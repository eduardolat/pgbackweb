-- name: BackupsServiceGetBackup :one
SELECT * FROM backups
WHERE id = @id;
