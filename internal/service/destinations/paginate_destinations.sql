-- name: DestinationsServicePaginateDestinationsCount :one
SELECT COUNT(*) FROM destinations;

-- name: DestinationsServicePaginateDestinations :many
SELECT
  *,
  pgp_sym_decrypt(access_key::bytea, @encryption_key) AS decrypted_access_key,
  pgp_sym_decrypt(secret_key::bytea, @encryption_key) AS decrypted_secret_key
FROM destinations
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
