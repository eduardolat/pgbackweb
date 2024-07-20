-- name: UsersServiceChangePassword :exec
UPDATE users
SET password = @password
WHERE id = @id;
