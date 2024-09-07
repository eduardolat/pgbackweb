-- name: DatabasesServiceGetDatabasesQty :one
SELECT 
  COUNT(*) AS all,
  COALESCE(SUM(CASE WHEN test_ok = true THEN 1 ELSE 0 END), 0)::INTEGER AS healthy,
  COALESCE(SUM(CASE WHEN test_ok = false THEN 1 ELSE 0 END), 0)::INTEGER AS unhealthy
FROM databases;
