package backups

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) ListBackups(
	ctx context.Context,
) ([]dbgen.Backup, error) {
	return s.dbgen.BackupsServiceListBackups(ctx)
}
