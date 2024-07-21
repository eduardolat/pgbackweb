-- name: DestinationsServiceGetDestination :one
SELECT * FROM destinations
WHERE id = @id;
