-- name: ExecutionsServiceGetDownloadLinkData :one
SELECT
  executions.path AS path,
  destinations.bucket_name AS bucket_name,
  destinations.region AS region,
  destinations.endpoint AS endpoint,
  pgp_sym_decrypt(destinations.access_key, sqlc.arg('decryption_key')::TEXT) AS decrypted_access_key,
  pgp_sym_decrypt(destinations.secret_key, sqlc.arg('decryption_key')::TEXT) AS decrypted_secret_key
FROM executions
JOIN backups ON backups.id = executions.backup_id
JOIN destinations ON destinations.id = backups.destination_id
WHERE executions.id = @execution_id;
