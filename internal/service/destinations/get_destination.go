package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) GetDestination(
	ctx context.Context, id uuid.UUID,
) (dbgen.DestinationsServiceGetDestinationRow, error) {
	return s.dbgen.DestinationsServiceGetDestination(
		ctx, dbgen.DestinationsServiceGetDestinationParams{
			ID:            id,
			EncryptionKey: *s.env.PBW_ENCRYPTION_KEY,
		},
	)
}
