package users

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) ChangePassword(
	ctx context.Context, params dbgen.UsersServiceChangePasswordParams,
) error {
	return s.dbgen.UsersServiceChangePassword(ctx, params)
}
