package executions

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) ListExecutions(
	ctx context.Context,
) ([]dbgen.Execution, error) {
	return s.dbgen.ExecutionsServiceListExecutions(ctx)
}
