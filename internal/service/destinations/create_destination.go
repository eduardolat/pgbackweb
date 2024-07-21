package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// CreateDestination creates a new destination entry.
func (s *Service) CreateDestination(
	ctx context.Context, params dbgen.DestinationsServiceCreateDestinationParams,
) (dbgen.Destination, error) {
	return s.dbgen.DestinationsServiceCreateDestination(ctx, params)
}
