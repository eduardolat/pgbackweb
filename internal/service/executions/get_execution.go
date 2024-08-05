package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) GetExecution(
	ctx context.Context, id uuid.UUID,
) (dbgen.ExecutionsServiceGetExecutionRow, error) {
	return s.dbgen.ExecutionsServiceGetExecution(ctx, id)
}
