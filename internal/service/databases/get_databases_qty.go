package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetDatabasesQty(
	ctx context.Context,
) (dbgen.DatabasesServiceGetDatabasesQtyRow, error) {
	return s.dbgen.DatabasesServiceGetDatabasesQty(ctx)
}
