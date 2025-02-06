package destinations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateDestination(
	ctx context.Context, params dbgen.DestinationsServiceUpdateDestinationParams,
) (dbgen.Destination, error) {
	err := s.TestDestination(
		params.AccessKey.String, params.SecretKey.String, params.Region.String,
		params.Endpoint.String, params.BucketName.String,
	)
	if err != nil {
		return dbgen.Destination{}, err
	}

	params.EncryptionKey = s.env.PBW_ENCRYPTION_KEY
	dest, err := s.dbgen.DestinationsServiceUpdateDestination(ctx, params)

	_ = s.TestDestinationAndStoreResult(ctx, dest.ID)

	return dest, err
}
