package users

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetUserByEmail(
	ctx context.Context, email string,
) (dbgen.User, error) {
	return s.dbgen.UsersServiceGetUserByEmail(ctx, email)
}
