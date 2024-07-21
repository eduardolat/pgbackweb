package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateExecution(
	ctx context.Context, params dbgen.ExecutionsServiceUpdateExecutionParams,
) (dbgen.Execution, error) {
	return s.dbgen.ExecutionsServiceUpdateExecution(ctx, params)
}
