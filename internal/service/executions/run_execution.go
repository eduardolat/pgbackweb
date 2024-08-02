package executions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration/postgres"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/google/uuid"
)

// RunExecution runs a backup execution
func (s *Service) RunExecution(ctx context.Context, backupID uuid.UUID) error {
	updateExec := func(params dbgen.ExecutionsServiceUpdateExecutionParams) error {
		_, err := s.dbgen.ExecutionsServiceUpdateExecution(
			ctx, params,
		)
		return err
	}

	logError := func(err error) {
		logger.Error("error running backup", logger.KV{
			"backup_id": backupID.String(),
			"error":     err.Error(),
		})
	}

	back, err := s.dbgen.ExecutionsServiceGetBackupData(
		ctx, dbgen.ExecutionsServiceGetBackupDataParams{
			BackupID:      backupID,
			EncryptionKey: *s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil {
		logError(err)
		return err
	}

	ex, err := s.CreateExecution(ctx, dbgen.ExecutionsServiceCreateExecutionParams{
		BackupID: backupID,
		Status:   "running",
	})
	if err != nil {
		logError(err)
		return err
	}

	err = s.ints.S3Client.Ping(
		back.DecryptedDestinationAccessKey, back.DecryptedDestinationSecretKey,
		back.DestinationRegion, back.DestinationEndpoint, back.DestinationBucketName,
	)
	if err != nil {
		logError(err)
		return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
			ID:         ex.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	pgVersion, err := s.ints.PGClient.ParseVersion(back.DatabasePgVersion)
	if err != nil {
		logError(err)
		return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
			ID:         ex.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	err = s.ints.PGClient.Ping(pgVersion, back.DecryptedDatabaseConnectionString)
	if err != nil {
		logError(err)
		return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
			ID:         ex.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	dumpBytes, err := s.ints.PGClient.DumpZip(
		pgVersion, back.DecryptedDatabaseConnectionString, postgres.DumpParams{
			DataOnly:   back.BackupOptDataOnly,
			SchemaOnly: back.BackupOptSchemaOnly,
			Clean:      back.BackupOptClean,
			IfExists:   back.BackupOptIfExists,
			Create:     back.BackupOptCreate,
			NoComments: back.BackupOptNoComments,
		},
	)
	if err != nil {
		logError(err)
		return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
			ID:         ex.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	date := time.Now().Format(timeutil.LayoutSlashYYYYMMDD)
	file := fmt.Sprintf(
		"dump-%s-%s.zip",
		time.Now().Format(timeutil.LayoutYYYYMMDDHHMMSS),
		uuid.NewString(),
	)
	path := strutil.CreatePath(false, back.BackupDestDir, date, file)

	_, err = s.ints.S3Client.Upload(
		back.DecryptedDestinationAccessKey, back.DecryptedDestinationSecretKey,
		back.DestinationRegion, back.DestinationEndpoint, back.DestinationBucketName,
		path, dumpBytes,
	)
	if err != nil {
		logError(err)
		return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
			ID:         ex.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			Path:       sql.NullString{Valid: true, String: path},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	logger.Info("backup created successfully", logger.KV{
		"backup_id":    backupID.String(),
		"execution_id": ex.ID.String(),
	})
	return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
		ID:         ex.ID,
		Status:     sql.NullString{Valid: true, String: "success"},
		Message:    sql.NullString{Valid: true, String: "Backup created successfully"},
		Path:       sql.NullString{Valid: true, String: path},
		FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
	})
}
