-- name: DatabasesServiceGetDatabase :one
SELECT
  *,
  pgp_sym_decrypt(connection_string::bytea, @encryption_key) AS decrypted_connection_string
FROM databases
WHERE id = @id;
