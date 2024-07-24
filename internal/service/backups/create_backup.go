package backups

import (
	"context"
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/validate"
)

func (s *Service) CreateBackup(
	ctx context.Context, params dbgen.BackupsServiceCreateBackupParams,
) (dbgen.Backup, error) {
	if !validate.CronExpression(params.CronExpression) {
		return dbgen.Backup{}, fmt.Errorf("invalid cron expression")
	}

	backup, err := s.dbgen.BackupsServiceCreateBackup(ctx, params)
	if err != nil {
		return backup, err
	}

	if !backup.IsActive {
		return backup, s.jobRemove(backup.ID)
	}

	return backup, s.jobUpsert(backup.ID, backup.TimeZone, backup.CronExpression)
}
