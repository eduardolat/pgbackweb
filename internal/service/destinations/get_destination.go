package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

// GetDestination retrieves a destination entry by ID.
func (s *Service) GetDestination(
	ctx context.Context, id uuid.UUID,
) (dbgen.Destination, error) {
	return s.dbgen.DestinationsServiceGetDestination(ctx, id)
}
