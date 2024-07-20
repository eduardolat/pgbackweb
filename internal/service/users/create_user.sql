-- name: UsersServiceCreateUser :one
INSERT INTO users (name, email, password)
VALUES (@name, lower(@email), @password)
RETURNING *;
