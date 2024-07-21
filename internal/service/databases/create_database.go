package databases

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateDatabase(
	ctx context.Context, params dbgen.DatabasesServiceCreateDatabaseParams,
) (dbgen.Database, error) {
	params.EncryptionKey = *s.env.PBW_ENCRYPTION_KEY
	return s.dbgen.DatabasesServiceCreateDatabase(ctx, params)
}
