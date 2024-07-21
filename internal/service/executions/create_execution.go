package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateExecution(
	ctx context.Context, params dbgen.ExecutionsServiceCreateExecutionParams,
) (dbgen.Execution, error) {
	return s.dbgen.ExecutionsServiceCreateExecution(ctx, params)
}
