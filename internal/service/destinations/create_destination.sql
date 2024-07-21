-- name: DestinationsServiceCreateDestination :one
INSERT INTO destinations (name, bucket_name, access_key, secret_key, region, endpoint)
VALUES (@name, @bucket_name, @access_key, @secret_key, @region, @endpoint)
RETURNING *;
