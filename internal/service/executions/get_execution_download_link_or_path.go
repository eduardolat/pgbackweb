package executions

import (
	"context"
	"fmt"
	"time"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

// GetExecutionDownloadLinkOrPath returns a download link for the file associated
// with the given execution. If the execution is stored locally, the link will
// be a file path.
//
// Returns a boolean indicating if the file is locally stored and the download
// link/path.
func (s *Service) GetExecutionDownloadLinkOrPath(
	ctx context.Context, executionID uuid.UUID,
) (bool, string, error) {
	data, err := s.dbgen.ExecutionsServiceGetDownloadLinkOrPathData(
		ctx, dbgen.ExecutionsServiceGetDownloadLinkOrPathDataParams{
			ExecutionID:   executionID,
			DecryptionKey: s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil {
		return false, "", err
	}

	if !data.Path.Valid {
		return false, "", fmt.Errorf("execution has no file associated")
	}

	if data.IsLocal {
		return true, s.ints.StorageClient.LocalGetFullPath(data.Path.String), nil
	}

	link, err := s.ints.StorageClient.S3GetDownloadLink(
		data.DecryptedAccessKey, data.DecryptedSecretKey, data.Region.String,
		data.Endpoint.String, data.BucketName.String, data.Path.String, time.Hour*12,
	)
	if err != nil {
		return false, "", err
	}
	return false, link, nil
}
