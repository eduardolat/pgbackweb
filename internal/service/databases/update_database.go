package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// UpdateDatabase updates an existing database entry.
func (s *Service) UpdateDatabase(
	ctx context.Context, params dbgen.DatabasesServiceUpdateDatabaseParams,
) (dbgen.Database, error) {
	err := s.TestDatabase(
		ctx, params.PgVersion.String, params.ConnectionString.String,
	)
	if err != nil {
		return dbgen.Database{}, err
	}

	params.EncryptionKey = *s.env.PBW_ENCRYPTION_KEY
	return s.dbgen.DatabasesServiceUpdateDatabase(ctx, params)
}
