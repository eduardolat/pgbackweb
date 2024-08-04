-- name: ExecutionsServiceGetExecutionForSoftDelete :one
SELECT
  executions.id as execution_id,
  executions.path as execution_path,

  backups.id as backup_id,
  backups.is_local as backup_is_local,

  destinations.bucket_name as destination_bucket_name,
  destinations.region as destination_region,
  destinations.endpoint as destination_endpoint,
  (
    CASE WHEN destinations.access_key IS NOT NULL
    THEN pgp_sym_decrypt(destinations.access_key, sqlc.arg('encryption_key')::TEXT)
    ELSE ''
    END
  ) AS decrypted_destination_access_key,
  (
    CASE WHEN destinations.secret_key IS NOT NULL
    THEN pgp_sym_decrypt(destinations.secret_key, sqlc.arg('encryption_key')::TEXT)
    ELSE ''
    END
  ) AS decrypted_destination_secret_key
FROM executions
INNER JOIN backups ON backups.id = executions.backup_id
LEFT JOIN destinations ON destinations.id = backups.destination_id
WHERE executions.id = @execution_id;

-- name: ExecutionsServiceSoftDeleteExecution :exec
UPDATE executions
SET
  status = 'deleted',
  deleted_at = NOW()
WHERE id = @id;
