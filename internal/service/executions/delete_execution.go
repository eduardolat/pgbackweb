package executions

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteExecution(
	ctx context.Context, id uuid.UUID,
) error {
	return s.dbgen.ExecutionsServiceDeleteExecution(ctx, id)
}
