package restorations

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/google/uuid"
)

// RunRestoration runs a backup restoration
func (s *Service) RunRestoration(
	ctx context.Context,
	executionID uuid.UUID,
	databaseID uuid.NullUUID,
	connString string,
) error {
	updateRes := func(params dbgen.RestorationsServiceUpdateRestorationParams) error {
		_, err := s.dbgen.RestorationsServiceUpdateRestoration(
			ctx, params,
		)
		return err
	}

	logError := func(err error) {
		dbID := "empty"
		if databaseID.Valid {
			dbID = databaseID.UUID.String()
		}
		logger.Error("error running restoration", logger.KV{
			"execution_id": executionID.String(),
			"database_id":  dbID,
			"error":        err.Error(),
		})
	}

	res, err := s.CreateRestoration(ctx, dbgen.RestorationsServiceCreateRestorationParams{
		ExecutionID: executionID,
		DatabaseID:  databaseID,
		Status:      "running",
	})
	if err != nil {
		logError(err)
		return err
	}

	if !databaseID.Valid && connString == "" {
		err := fmt.Errorf("database_id or connection_string must be provided")
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	execution, err := s.executionsService.GetExecution(ctx, executionID)
	if err != nil {
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	if execution.Status != "success" || !execution.Path.Valid {
		err := fmt.Errorf("backup execution must be successful")
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	if databaseID.Valid {
		db, err := s.databasesService.GetDatabase(ctx, databaseID.UUID)
		if err != nil {
			logError(err)
			return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
				ID:         res.ID,
				Status:     sql.NullString{Valid: true, String: "failed"},
				Message:    sql.NullString{Valid: true, String: err.Error()},
				FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
			})
		}
		connString = db.DecryptedConnectionString
	}

	pgVersion, err := s.ints.PGClient.ParseVersion(execution.DatabasePgVersion)
	if err != nil {
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	err = s.ints.PGClient.Ping(pgVersion, connString)
	if err != nil {
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	isLocal, zipURLOrPath, err := s.executionsService.GetExecutionDownloadLinkOrPath(
		ctx, executionID,
	)
	if err != nil {
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	err = s.ints.PGClient.RestoreZip(
		pgVersion, connString, isLocal, zipURLOrPath,
	)
	if err != nil {
		logError(err)
		return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
			ID:         res.ID,
			Status:     sql.NullString{Valid: true, String: "failed"},
			Message:    sql.NullString{Valid: true, String: err.Error()},
			FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
		})
	}

	logger.Info("backup restored successfully", logger.KV{
		"restoration_id": res.ID.String(),
		"execution_id":   executionID.String(),
	})
	return updateRes(dbgen.RestorationsServiceUpdateRestorationParams{
		ID:         res.ID,
		Status:     sql.NullString{Valid: true, String: "success"},
		Message:    sql.NullString{Valid: true, String: "Backup restored successfully"},
		FinishedAt: sql.NullTime{Valid: true, Time: time.Now()},
	})
}
