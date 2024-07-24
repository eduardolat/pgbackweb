-- name: DatabasesServiceUpdateDatabase :one
UPDATE databases
SET
  name = COALESCE(sqlc.narg('name'), name),
  pg_version = COALESCE(sqlc.narg('pg_version'), pg_version),
  connection_string = CASE
    WHEN sqlc.narg('connection_string')::TEXT IS NOT NULL
    THEN pgp_sym_encrypt(
      sqlc.narg('connection_string')::TEXT, sqlc.arg('encryption_key')::TEXT
    )
    ELSE connection_string
  END
WHERE id = @id
RETURNING *;
