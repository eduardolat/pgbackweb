package executions

import (
	"context"
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) GetExecutionDownloadLink(
	ctx context.Context, executionID uuid.UUID,
) (string, error) {
	data, err := s.dbgen.ExecutionsServiceGetDownloadLinkData(
		ctx, dbgen.ExecutionsServiceGetDownloadLinkDataParams{
			ExecutionID:   executionID,
			DecryptionKey: *s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil {
		return "", err
	}

	if !data.Path.Valid {
		return "", fmt.Errorf("execution has no file associated")
	}

	return s.ints.S3Client.GetDownloadLink(
		data.DecryptedAccessKey, data.DecryptedSecretKey, data.Region,
		data.Endpoint, data.BucketName, data.Path.String, time.Hour*12,
	)
}
