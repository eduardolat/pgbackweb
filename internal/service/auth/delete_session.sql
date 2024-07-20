-- name: AuthServiceDeleteSession :exec
DELETE FROM sessions WHERE id = @id;
