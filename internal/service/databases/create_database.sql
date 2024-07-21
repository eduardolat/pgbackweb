-- name: DatabasesServiceCreateDatabase :one
INSERT INTO databases (name, connection_string)
VALUES (@name, @connection_string)
RETURNING *;
