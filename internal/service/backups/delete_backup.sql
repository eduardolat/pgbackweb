-- name: BackupsServiceDeleteBackup :exec
DELETE FROM backups
WHERE id = @id;
