package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// ListDatabases lists all database entries.
func (s *Service) ListDatabases(
	ctx context.Context,
) ([]dbgen.Database, error) {
	return s.dbgen.DatabasesServiceListDatabases(ctx)
}
