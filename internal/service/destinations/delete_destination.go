package destinations

import (
	"context"

	"github.com/google/uuid"
)

// DeleteDestination deletes a destination entry by ID.
func (s *Service) DeleteDestination(
	ctx context.Context, id uuid.UUID,
) error {
	return s.dbgen.DestinationsServiceDeleteDestination(ctx, id)
}
