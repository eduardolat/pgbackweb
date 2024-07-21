-- name: DatabasesServiceGetDatabase :one
SELECT * FROM databases
WHERE id = @id;
