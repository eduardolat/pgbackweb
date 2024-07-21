-- name: ExecutionsServiceGetBackupData :one
SELECT
  backups.is_active as backup_is_active,
  backups.dest_dir as backup_dest_dir,
  backups.opt_data_only as backup_opt_data_only,
  backups.opt_schema_only as backup_opt_schema_only,
  backups.opt_clean as backup_opt_clean,
  backups.opt_if_exists as backup_opt_if_exists,
  backups.opt_create as backup_opt_create,	
  backups.opt_no_comments as backup_opt_no_comments,

  pgp_sym_decrypt(
    databases.connection_string::bytea, @encryption_key
  ) AS database_connection_string,
  databases.pg_version as database_pg_version,

  destinations.bucket_name as destination_bucket_name,
  destinations.region as destination_region,
  destinations.endpoint as destination_endpoint,
  pgp_sym_decrypt(
    destinations.access_key, @encryption_key
  ) AS destination_access_key,
  pgp_sym_decrypt(
    destinations.secret_key, @encryption_key
  ) AS destination_secret_key
FROM backups
JOIN databases ON backups.database_id = databases.id
JOIN destinations ON backups.destination_id = destinations.id
WHERE backups.id = @backup_id;
