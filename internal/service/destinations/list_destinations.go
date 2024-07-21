package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

// ListDestinations lists all destination entries.
func (s *Service) ListDestinations(
	ctx context.Context,
) ([]dbgen.Destination, error) {
	return s.dbgen.DestinationsServiceListDestinations(ctx)
}
