-- name: ExecutionsServiceGetExecutionForSoftDelete :one
SELECT
  executions.id as execution_id,
  executions.path as execution_path,

  backups.id as backup_id,

  destinations.bucket_name as destination_bucket_name,
  destinations.region as destination_region,
  destinations.endpoint as destination_endpoint,
  pgp_sym_decrypt(
    destinations.access_key::bytea, @encryption_key
  ) AS destination_access_key,
  pgp_sym_decrypt(
    destinations.secret_key::bytea, @encryption_key
  ) AS destination_secret_key
FROM executions
JOIN backups ON backups.id = executions.backup_id
JOIN destinations ON destinations.id = backups.destination_id
WHERE executions.id = @execution_id;

-- name: ExecutionsServiceSoftDeleteExecution :exec
UPDATE executions
SET
  status = 'deleted',
  deleted_at = NOW()
WHERE id = @id;
