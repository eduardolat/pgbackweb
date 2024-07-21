package backups

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) ToggleIsActive(ctx context.Context, backupID uuid.UUID) error {
	backup, err := s.dbgen.BackupsServiceToggleIsActive(ctx, backupID)
	if err != nil {
		return err
	}

	if !backup.IsActive {
		return s.jobRemove(backupID)
	}

	return s.jobUpsert(backupID, backup.TimeZone, backup.CronExpression)
}
