package restorations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetRestorationsQty(
	ctx context.Context,
) (dbgen.RestorationsServiceGetRestorationsQtyRow, error) {
	return s.dbgen.RestorationsServiceGetRestorationsQty(ctx)
}
