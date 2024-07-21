-- name: DestinationsServiceUpdateDestination :one
UPDATE destinations
SET
  name = @name,
  bucket_name = @bucket_name,
  region = @region,
  endpoint = @endpoint,
  access_key = pgp_sym_encrypt(@access_key, @encryption_key),
  secret_key = pgp_sym_encrypt(@secret_key, @encryption_key)
WHERE id = @id
RETURNING *;
