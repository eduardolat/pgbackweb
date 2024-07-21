-- name: DatabasesServiceUpdateDatabase :one
UPDATE databases
SET
  name = @name,
  connection_string = pgp_sym_encrypt(@connection_string, @encryption_key)
WHERE id = @id
RETURNING *;
