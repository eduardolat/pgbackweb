-- name: DestinationsServiceGetAllDestinations :many
SELECT
  *,
  pgp_sym_decrypt(access_key, @encryption_key) AS decrypted_access_key,
  pgp_sym_decrypt(secret_key, @encryption_key) AS decrypted_secret_key
FROM destinations
ORDER BY created_at DESC;
