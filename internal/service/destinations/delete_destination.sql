-- name: DestinationsServiceDeleteDestination :exec
DELETE FROM destinations
WHERE id = @id;
