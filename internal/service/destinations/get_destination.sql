-- name: DestinationsServiceGetDestination :one
SELECT
  *,
  pgp_sym_decrypt(access_key, @encryption_key) AS decrypted_access_key,
  pgp_sym_decrypt(secret_key, @encryption_key) AS decrypted_secret_key
FROM destinations
WHERE id = @id;
