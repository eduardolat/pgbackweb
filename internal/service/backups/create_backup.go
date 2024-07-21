package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateBackup(
	ctx context.Context, params dbgen.BackupsServiceCreateBackupParams,
) (dbgen.Backup, error) {
	backup, err := s.dbgen.BackupsServiceCreateBackup(ctx, params)
	if err != nil {
		return backup, err
	}

	if !backup.IsActive {
		return backup, s.jobRemove(backup.ID)
	}

	return backup, s.jobUpsert(backup.ID, backup.TimeZone, backup.CronExpression)
}
