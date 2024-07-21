-- name: DatabasesServiceCreateDatabase :one
INSERT INTO databases (
  name, connection_string, pg_version
)
VALUES (
  @name, pgp_sym_encrypt(@connection_string, @encryption_key), @pg_version
)
RETURNING *;
