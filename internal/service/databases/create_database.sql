-- name: DatabasesServiceCreateDatabase :one
INSERT INTO databases (name, connection_string)
VALUES (@name, pgp_sym_encrypt(@connection_string, @encryption_key))
RETURNING *;
