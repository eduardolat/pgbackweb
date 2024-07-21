package backups

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteBackup(
	ctx context.Context, id uuid.UUID,
) error {
	err := s.jobRemove(id)
	if err != nil {
		return err
	}

	return s.dbgen.BackupsServiceDeleteBackup(ctx, id)
}
