package executions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration/postgres"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/metrics"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
	"github.com/eduardolat/pgbackweb/internal/util/timeutil"
	"github.com/google/uuid"
)

// RunExecution runs a backup execution
func (s *Service) RunExecution(ctx context.Context, backupID uuid.UUID) error {
	startTime := time.Now()

	logError := func(err error) {
		logger.Error("error running backup", logger.KV{
			"backup_id": backupID.String(),
			"error":     err.Error(),
		})
	}

	back, err := s.dbgen.ExecutionsServiceGetBackupData(
		ctx, dbgen.ExecutionsServiceGetBackupDataParams{
			BackupID:      backupID,
			EncryptionKey: s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil {
		logError(err)
		return err
	}

	metrics.BackupsRunning.WithLabelValues(back.BackupName, back.DatabaseName).Inc()
	defer metrics.BackupsRunning.WithLabelValues(back.BackupName, back.DatabaseName).Dec()

	ex, err := s.CreateExecution(ctx, dbgen.ExecutionsServiceCreateExecutionParams{
		BackupID: backupID,
		Status:   "running",
	})
	if err != nil {
		logError(err)
		return err
	}

	logger.Info("execution started", logger.KV{
		"backup_id":     backupID.String(),
		"execution_id":  ex.ID.String(),
		"backup_name":   back.BackupName,
		"database_name": back.DatabaseName,
		"destination":   back.DestinationName.String,
	})

	updateExec := func(params dbgen.ExecutionsServiceUpdateExecutionParams) error {
		// Update database first
		_, err := s.dbgen.ExecutionsServiceUpdateExecution(ctx, params)
		if err != nil {
			return err
		}

		// Only emit side effects if database update succeeded
		if params.Status.String == "success" {
			s.webhooksService.RunExecutionSuccess(backupID)
			metrics.BackupsTotal.WithLabelValues("success", back.BackupName, back.DatabaseName).Inc()
			metrics.LastBackupDuration.WithLabelValues(back.BackupName, back.DatabaseName).Set(time.Since(startTime).Seconds())
			metrics.LastBackupStatus.WithLabelValues(back.BackupName, back.DatabaseName).Set(1.0)
			logger.Info("execution completed successfully", logger.KV{
				"backup_id":     backupID.String(),
				"execution_id":  ex.ID.String(),
				"backup_name":   back.BackupName,
				"database_name": back.DatabaseName,
				"duration_sec":  time.Since(startTime).Seconds(),
				"file_size":     params.FileSize.Int64,
				"path":          params.Path.String,
			})
		}

		if params.Status.String == "failed" {
			s.webhooksService.RunExecutionFailed(backupID)
			metrics.BackupsTotal.WithLabelValues("failed", back.BackupName, back.DatabaseName).Inc()
			metrics.LastBackupDuration.WithLabelValues(back.BackupName, back.DatabaseName).Set(time.Since(startTime).Seconds())
			metrics.LastBackupStatus.WithLabelValues(back.BackupName, back.DatabaseName).Set(0.0)
			logger.Error("execution failed", logger.KV{
				"backup_id":     backupID.String(),
				"execution_id":  ex.ID.String(),
				"backup_name":   back.BackupName,
				"database_name": back.DatabaseName,
				"error_message": params.Message.String,
				"duration_sec":  time.Since(startTime).Seconds(),
			})
		}

		return nil
	}

	if !back.BackupIsLocal {
		err = s.ints.StorageClient.S3Test(
			back.DecryptedDestinationAccessKey, back.DecryptedDestinationSecretKey,
			back.DestinationRegion.String, back.DestinationEndpoint.String,
			back.DestinationBucketName.String,
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

	err = s.ints.PGClient.Test(pgVersion, back.DecryptedDatabaseConnectionString)
	if err != nil {
		logError(err)
		return updateExec(dbgen.ExecutionsServiceUpdateExecutionParams{
			ID:         ex.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	dumpReader := s.ints.PGClient.DumpZip(
		pgVersion, back.DecryptedDatabaseConnectionString, postgres.DumpParams{
			DataOnly:   back.BackupOptDataOnly,
			SchemaOnly: back.BackupOptSchemaOnly,
			Clean:      back.BackupOptClean,
			IfExists:   back.BackupOptIfExists,
			Create:     back.BackupOptCreate,
			NoComments: back.BackupOptNoComments,
		},
	)

	date := time.Now().Format(timeutil.LayoutSlashYYYYMMDD)
	file := fmt.Sprintf(
		"dump-%s-%s.zip",
		time.Now().Format(timeutil.LayoutYYYYMMDDHHMMSS),
		uuid.NewString(),
	)
	path := strutil.CreatePath(false, back.BackupDestDir, date, file)
	fileSize := int64(0)

	if back.BackupIsLocal {
		fileSize, err = s.ints.StorageClient.LocalUpload(path, dumpReader)
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
	}

	if !back.BackupIsLocal {
		fileSize, err = s.ints.StorageClient.S3Upload(
			back.DecryptedDestinationAccessKey, back.DecryptedDestinationSecretKey,
			back.DestinationRegion.String, back.DestinationEndpoint.String,
			back.DestinationBucketName.String, path, dumpReader,
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
		FileSize:   sql.NullInt64{Valid: true, Int64: fileSize},
	})
}
