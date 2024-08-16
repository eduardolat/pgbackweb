package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateDatabase(
	ctx context.Context, params dbgen.DatabasesServiceCreateDatabaseParams,
) (dbgen.Database, error) {
	err := s.TestDatabase(ctx, params.PgVersion, params.ConnectionString)
	if err != nil {
		return dbgen.Database{}, err
	}

	params.EncryptionKey = *s.env.PBW_ENCRYPTION_KEY
	db, err := s.dbgen.DatabasesServiceCreateDatabase(ctx, params)

	_ = s.TestDatabaseAndStoreResult(ctx, db.ID)

	return db, err
}
