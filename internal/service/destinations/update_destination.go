package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateDestination(
	ctx context.Context, params dbgen.DestinationsServiceUpdateDestinationParams,
) (dbgen.Destination, error) {
	return s.dbgen.DestinationsServiceUpdateDestination(ctx, params)
}
