-- name: DestinationsServiceCreateDestination :one
INSERT INTO destinations (
  name, bucket_name, region, endpoint,
  access_key, secret_key
)
VALUES (
  @name, @bucket_name, @region, @endpoint,
  pgp_sym_encrypt(@access_key, @encryption_key),
  pgp_sym_encrypt(@secret_key, @encryption_key)
)
RETURNING *;
