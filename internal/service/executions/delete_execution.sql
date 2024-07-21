-- name: ExecutionsServiceDeleteExecution :exec
DELETE FROM executions
WHERE id = @id;
