-- name: DatabasesServiceUpdateDatabase :one
UPDATE databases
SET
  name = @name,
  connection_string = pgp_sym_encrypt(@connection_string, @encryption_key),
  pg_version = @pg_version
WHERE id = @id
RETURNING *;
