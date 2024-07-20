-- name: UsersServiceUpdateUser :one
UPDATE users
SET
  name = @name,
  email = lower(@email),
  password = @password
WHERE id = @id
RETURNING *;
