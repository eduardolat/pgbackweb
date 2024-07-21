package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateBackup(
	ctx context.Context, params dbgen.BackupsServiceUpdateBackupParams,
) (dbgen.Backup, error) {
	backup, err := s.dbgen.BackupsServiceUpdateBackup(ctx, params)
	if err != nil {
		return backup, err
	}

	if !backup.IsActive {
		return backup, s.jobRemove(backup.ID)
	}

	return backup, s.jobUpsert(backup.ID, backup.TimeZone, backup.CronExpression)
}
