package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateDestination(
	ctx context.Context, params dbgen.DestinationsServiceCreateDestinationParams,
) (dbgen.Destination, error) {
	err := s.TestDestination(
		params.AccessKey, params.SecretKey, params.Region, params.Endpoint,
		params.BucketName,
	)
	if err != nil {
		return dbgen.Destination{}, err
	}

	params.EncryptionKey = *s.env.PBW_ENCRYPTION_KEY
	dest, err := s.dbgen.DestinationsServiceCreateDestination(ctx, params)

	_ = s.TestDestinationAndStoreResult(ctx, dest.ID)

	return dest, err
}
