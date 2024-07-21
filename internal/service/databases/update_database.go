package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// UpdateDatabase updates an existing database entry.
func (s *Service) UpdateDatabase(
	ctx context.Context, params dbgen.DatabasesServiceUpdateDatabaseParams,
) (dbgen.Database, error) {
	params.EncryptionKey = *s.env.PBW_ENCRYPTION_KEY
	return s.dbgen.DatabasesServiceUpdateDatabase(ctx, params)
}
