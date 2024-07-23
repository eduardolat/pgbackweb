-- name: DatabasesServicePaginateDatabasesCount :one
SELECT COUNT(*) FROM databases;

-- name: DatabasesServicePaginateDatabases :many
SELECT
  *,
  pgp_sym_decrypt(connection_string::bytea, @encryption_key) AS decrypted_connection_string
FROM databases
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
