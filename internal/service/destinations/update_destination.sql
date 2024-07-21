-- name: DestinationsServiceUpdateDestination :one
UPDATE destinations
SET
  name = @name,
  bucket_name = @bucket_name,
  access_key = @access_key,
  secret_key = @secret_key,
  region = @region,
  endpoint = @endpoint
WHERE id = @id
RETURNING *;
