-- name: DatabasesServiceGetAllDatabases :one
SELECT
  *,
  pgp_sym_decrypt(connection_string, @encryption_key) AS decrypted_connection_string
FROM databases
ORDER BY created_at DESC;
