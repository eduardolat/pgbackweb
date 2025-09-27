package users

import (
	"context"
	"database/sql"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/util/cryptoutil"
)

func (s *Service) CreateUser(
	ctx context.Context, params dbgen.UsersServiceCreateUserParams,
) (dbgen.User, error) {
	// Convert sql.NullString to string for hashing
	passwordStr := ""
	if params.Password.Valid {
		passwordStr = params.Password.String
	}

	hash, err := cryptoutil.CreateBcryptHash(passwordStr)
	if err != nil {
		return dbgen.User{}, err
	}

	// Convert hash back to sql.NullString
	params.Password = sql.NullString{String: hash, Valid: true}

	return s.dbgen.UsersServiceCreateUser(ctx, params)
}
