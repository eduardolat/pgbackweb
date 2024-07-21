package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// UpdateDestination updates an existing destination entry.
func (s *Service) UpdateDestination(
	ctx context.Context, params dbgen.DestinationsServiceUpdateDestinationParams,
) (dbgen.Destination, error) {
	return s.dbgen.DestinationsServiceUpdateDestination(ctx, params)
}
