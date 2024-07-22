package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/logger"
)

func (s *Service) ScheduleAll() error {
	activeBackups, err := s.dbgen.BackupsServiceGetScheduleAllData(
		context.Background(),
	)
	if err != nil {
		return err
	}

	for _, backup := range activeBackups {
		if !backup.IsActive {
			err := s.jobRemove(backup.ID)
			if err != nil {
				return err
			}
		}

		if backup.IsActive {
			err := s.jobUpsert(backup.ID, backup.TimeZone, backup.CronExpression)
			if err != nil {
				return err
			}
		}
	}

	logger.Info("all active backups scheduled")

	return nil
}
