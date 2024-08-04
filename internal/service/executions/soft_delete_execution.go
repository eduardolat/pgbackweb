package executions

import (
	"context"
	"database/sql"
	"errors"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) SoftDeleteExecution(
	ctx context.Context, executionID uuid.UUID,
) error {
	execution, err := s.dbgen.ExecutionsServiceGetExecutionForSoftDelete(
		ctx, dbgen.ExecutionsServiceGetExecutionForSoftDeleteParams{
			ExecutionID:   executionID,
			EncryptionKey: *s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}

	if execution.ExecutionPath.Valid && !execution.BackupIsLocal {
		err := s.ints.StorageClient.S3Delete(
			execution.DecryptedDestinationAccessKey, execution.DecryptedDestinationSecretKey,
			execution.DestinationRegion.String, execution.DestinationEndpoint.String,
			execution.DestinationBucketName.String, execution.ExecutionPath.String,
		)
		if err != nil {
			return err
		}
	}

	if execution.ExecutionPath.Valid && execution.BackupIsLocal {
		err := s.ints.StorageClient.LocalDelete(execution.ExecutionPath.String)
		if err != nil {
			return err
		}
	}

	return s.dbgen.ExecutionsServiceSoftDeleteExecution(ctx, executionID)
}
