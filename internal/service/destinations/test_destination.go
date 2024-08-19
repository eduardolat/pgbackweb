package destinations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) TestDestinationAndStoreResult(
	ctx context.Context, destinationID uuid.UUID,
) error {
	storeRes := func(ok bool, err error) error {
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}

		secondErr := s.dbgen.DestinationsServiceSetTestData(
			ctx, dbgen.DestinationsServiceSetTestDataParams{
				DestinationID: destinationID,
				TestOk:        sql.NullBool{Valid: true, Bool: ok},
				TestError:     sql.NullString{Valid: true, String: errMsg},
			},
		)
		if secondErr != nil {
			return fmt.Errorf("error storing test result: %w: %w", secondErr, err)
		}
		return err
	}

	dest, err := s.GetDestination(ctx, destinationID)
	if err != nil {
		return storeRes(false, fmt.Errorf("error getting destination: %w", err))
	}

	err = s.TestDestination(
		dest.DecryptedAccessKey, dest.DecryptedSecretKey, dest.Region,
		dest.Endpoint, dest.BucketName,
	)
	if err != nil && dest.TestOk.Valid && dest.TestOk.Bool {
		s.webhooksService.RunDestinationUnhealthy(dest.ID)
	}
	if err != nil {
		return storeRes(false, err)
	}

	if dest.TestOk.Valid && !dest.TestOk.Bool {
		s.webhooksService.RunDestinationHealthy(dest.ID)
	}
	return storeRes(true, nil)
}

func (s *Service) TestDestination(
	accessKey, secretKey, region, endpoint, bucketName string,
) error {
	err := s.ints.StorageClient.S3Test(accessKey, secretKey, region, endpoint, bucketName)
	if err != nil {
		return fmt.Errorf("error testing destination: %w", err)
	}

	return nil
}
