package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetAllDatabases(
	ctx context.Context,
) ([]dbgen.DatabasesServiceGetAllDatabasesRow, error) {
	return s.dbgen.DatabasesServiceGetAllDatabases(
		ctx, *s.env.PBW_ENCRYPTION_KEY,
	)
}
