package users

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) UpdateUser(
	ctx context.Context, params dbgen.UsersServiceUpdateUserParams,
) (dbgen.User, error) {
	return s.dbgen.UsersServiceUpdateUser(ctx, params)
}
