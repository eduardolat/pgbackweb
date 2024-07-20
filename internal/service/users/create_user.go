package users

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) CreateUser(
	ctx context.Context, params dbgen.UsersServiceCreateUserParams,
) (dbgen.User, error) {
	return s.dbgen.UsersServiceCreateUser(ctx, params)
}
