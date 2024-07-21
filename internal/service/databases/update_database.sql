-- name: DatabasesServiceUpdateDatabase :one
UPDATE databases
SET
  name = @name,
  connection_string = @connection_string
WHERE id = @id
RETURNING *;
