package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateBackup(
	ctx context.Context, params dbgen.BackupsServiceUpdateBackupParams,
) (dbgen.Backup, error) {
	return s.dbgen.BackupsServiceUpdateBackup(ctx, params)
}
