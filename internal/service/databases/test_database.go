package databases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
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
	if err != nil {
		return storeRes(false, err)
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
