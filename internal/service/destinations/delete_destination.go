package destinations

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteDestination(
	ctx context.Context, id uuid.UUID,
) error {
	return s.dbgen.DestinationsServiceDeleteDestination(ctx, id)
}
