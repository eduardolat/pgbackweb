package restorations

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateRestoration(
	ctx context.Context, params dbgen.RestorationsServiceUpdateRestorationParams,
) (dbgen.Restoration, error) {
	return s.dbgen.RestorationsServiceUpdateRestoration(ctx, params)
}
