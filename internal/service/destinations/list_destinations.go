package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) ListDestinations(
	ctx context.Context,
) ([]dbgen.DestinationsServiceListDestinationsRow, error) {
	return s.dbgen.DestinationsServiceListDestinations(
		ctx, *s.env.PBW_ENCRYPTION_KEY,
	)
}
