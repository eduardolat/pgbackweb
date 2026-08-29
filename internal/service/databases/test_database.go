package databases

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
	"github.com/google/uuid"
)

func (s *Service) TestDatabaseAndStoreResult(
	ctx context.Context, databaseID uuid.UUID,
) error {
	storeRes := func(ok bool, err error) error {
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}

		secondErr := s.dbgen.DatabasesServiceSetTestData(
			ctx, dbgen.DatabasesServiceSetTestDataParams{
				DatabaseID: databaseID,
				TestOk:     sql.NullBool{Valid: true, Bool: ok},
				TestError:  sql.NullString{Valid: true, String: errMsg},
			},
		)
		if secondErr != nil {
			return fmt.Errorf("error storing test result: %w: %w", secondErr, err)
		}
		return err
	}

	db, err := s.GetDatabase(ctx, databaseID)
	if err != nil {
		return storeRes(false, fmt.Errorf("error getting database: %w", err))
	}

	err = s.TestDatabase(ctx, db.PgVersion, db.DecryptedConnectionString)
	buildDatabaseStatus := func(ok bool, testErr error) webhooks.DatabaseStatus {
		var errMsg string
		if testErr != nil {
			errMsg = testErr.Error()
		}

		var lastTestOk *bool
		if db.TestOk.Valid {
			val := db.TestOk.Bool
			lastTestOk = &val
		}

		var lastTestError *string
		if db.TestError.Valid {
			val := db.TestError.String
			lastTestError = &val
		}

		var lastTestAt *time.Time
		if db.LastTestAt.Valid {
			val := db.LastTestAt.Time
			lastTestAt = &val
		}
		info, err := webhooks.ParsePostgresURL(db.DecryptedConnectionString)
		if err != nil {
			logger.Error("Error parsing URL: ", logger.KV{"error": err})
		}

		return webhooks.DatabaseStatus{
			ID:               db.ID,
			Name:             db.Name,
			PgVersion:        db.PgVersion,
			User:             info.User,
			Host:             info.Host,
			Port:             info.Port,
			DBName:           info.DBName,
			Healthy:          ok,
			Error:            errMsg,
			Timestamp:        time.Now().UTC(),
			LastCheckedAt:    lastTestAt,
			LastErrorMessage: lastTestError,
			LastSuccess:      lastTestOk,
		}
	}

	if err != nil && db.TestOk.Valid && db.TestOk.Bool {
		databaseStatus := buildDatabaseStatus(false, err)

		s.webhooksService.RunDatabaseUnhealthy(db.ID, databaseStatus)
	}
	if err != nil {
		return storeRes(false, err)
	}

	if db.TestOk.Valid && !db.TestOk.Bool {
		databaseStatus := buildDatabaseStatus(true, err)

		s.webhooksService.RunDatabaseHealthy(db.ID, databaseStatus)
	}
	return storeRes(true, nil)
}

func (s *Service) TestDatabase(
	ctx context.Context, version, connString string,
) error {
	pgVersion, err := s.ints.PGClient.ParseVersion(version)
	if err != nil {
		return fmt.Errorf("error parsing PostgreSQL version: %w", err)
	}

	err = s.ints.PGClient.Test(pgVersion, connString)
	if err != nil {
		return fmt.Errorf("error testing database: %w", err)
	}

	return nil
}
