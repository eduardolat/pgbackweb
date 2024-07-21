package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateBackup(
	ctx context.Context, params dbgen.BackupsServiceCreateBackupParams,
) (dbgen.Backup, error) {
	return s.dbgen.BackupsServiceCreateBackup(ctx, params)
}
