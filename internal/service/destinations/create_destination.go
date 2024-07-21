package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateDestination(
	ctx context.Context, params dbgen.DestinationsServiceCreateDestinationParams,
) (dbgen.Destination, error) {
	params.EncryptionKey = *s.env.PBW_ENCRYPTION_KEY
	return s.dbgen.DestinationsServiceCreateDestination(ctx, params)
}
