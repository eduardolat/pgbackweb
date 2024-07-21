package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) ListBackupExecutions(
	ctx context.Context, backupID uuid.UUID,
) ([]dbgen.Execution, error) {
	return s.dbgen.ExecutionsServiceListBackupExecutions(ctx, backupID)
}
