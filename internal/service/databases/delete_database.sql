-- name: DatabasesServiceDeleteDatabase :exec
DELETE FROM databases
WHERE id = @id;
