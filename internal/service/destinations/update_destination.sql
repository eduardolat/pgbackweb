-- name: DestinationsServiceUpdateDestination :one
UPDATE destinations
SET
  name = COALESCE(sqlc.narg('name'), name),
  bucket_name = COALESCE(sqlc.narg('bucket_name'), bucket_name),
  region = COALESCE(sqlc.narg('region'), region),
  endpoint = COALESCE(sqlc.narg('endpoint'), endpoint),
  access_key = CASE
    WHEN sqlc.narg('access_key')::TEXT IS NOT NULL
    THEN pgp_sym_encrypt(sqlc.narg('access_key')::TEXT, sqlc.arg('encryption_key')::TEXT)
    ELSE access_key
  END,
  secret_key = CASE
    WHEN sqlc.narg('secret_key')::TEXT IS NOT NULL
    THEN pgp_sym_encrypt(sqlc.narg('secret_key')::TEXT, sqlc.arg('encryption_key')::TEXT)
    ELSE secret_key
  END
WHERE id = @id
RETURNING *;
