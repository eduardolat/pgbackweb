-- name: DestinationsServiceGetDestination :one
SELECT
  *,
  pgp_sym_decrypt(access_key, @encryption_key) AS access_key,
  pgp_sym_decrypt(secret_key, @encryption_key) AS secret_key
FROM destinations
WHERE id = @id;
