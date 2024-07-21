package databases

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteDatabase(
	ctx context.Context, id uuid.UUID,
) error {
	return s.dbgen.DatabasesServiceDeleteDatabase(ctx, id)
}
