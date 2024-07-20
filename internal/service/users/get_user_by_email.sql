-- name: UsersServiceGetUserByEmail :one
SELECT * FROM users WHERE email = @email;
