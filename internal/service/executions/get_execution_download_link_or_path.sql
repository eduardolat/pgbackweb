-- name: ExecutionsServiceGetDownloadLinkOrPathData :one
SELECT
  executions.path AS path,
  backups.is_local AS is_local,
  destinations.bucket_name AS bucket_name,
  destinations.region AS region,
  destinations.endpoint AS endpoint,
  destinations.endpoint as destination_endpoint,
  (
    CASE WHEN destinations.access_key IS NOT NULL
    THEN pgp_sym_decrypt(destinations.access_key, sqlc.arg('decryption_key')::TEXT)
    ELSE ''
    END
  ) AS decrypted_access_key,
  (
    CASE WHEN destinations.secret_key IS NOT NULL
    THEN pgp_sym_decrypt(destinations.secret_key, sqlc.arg('decryption_key')::TEXT)
    ELSE ''
    END
  ) AS decrypted_secret_key
FROM executions
INNER JOIN backups ON backups.id = executions.backup_id
LEFT JOIN destinations ON destinations.id = backups.destination_id
WHERE executions.id = @execution_id;
