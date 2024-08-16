-- name: DestinationsServiceGetDestinationsQty :one
SELECT 
  COUNT(*) AS all,
  SUM(CASE WHEN test_ok = true THEN 1 ELSE 0 END) AS healthy,
  SUM(CASE WHEN test_ok = false THEN 1 ELSE 0 END) AS unhealthy
FROM destinations;
