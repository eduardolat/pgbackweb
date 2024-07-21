package backups

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteBackup(
	ctx context.Context, id uuid.UUID,
) error {
	return s.dbgen.BackupsServiceDeleteBackup(ctx, id)
}
