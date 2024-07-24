package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) GetAllDestinations(
	ctx context.Context, id uuid.UUID,
) (dbgen.DestinationsServiceGetAllDestinationsRow, error) {
	return s.dbgen.DestinationsServiceGetAllDestinations(
		ctx, *s.env.PBW_ENCRYPTION_KEY,
	)
}
