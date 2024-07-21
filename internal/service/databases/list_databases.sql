-- name: DatabasesServiceListDatabases :many
SELECT
  *,
  pgp_sym_decrypt(connection_string::bytea, @encryption_key) AS connection_string
FROM databases;
