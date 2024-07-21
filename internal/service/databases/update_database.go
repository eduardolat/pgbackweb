package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// UpdateDatabase updates an existing database entry.
func (s *Service) UpdateDatabase(
	ctx context.Context, params dbgen.DatabasesServiceUpdateDatabaseParams,
) (dbgen.Database, error) {
	return s.dbgen.DatabasesServiceUpdateDatabase(ctx, params)
}
