package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetDestinationsQty(
	ctx context.Context,
) (dbgen.DestinationsServiceGetDestinationsQtyRow, error) {
	return s.dbgen.DestinationsServiceGetDestinationsQty(ctx)
}
