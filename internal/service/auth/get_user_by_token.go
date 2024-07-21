package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

func (s *Service) GetUserByToken(
	ctx context.Context, token string,
) (bool, dbgen.AuthServiceGetUserByTokenRow, error) {
	user, err := s.dbgen.AuthServiceGetUserByToken(
		ctx, dbgen.AuthServiceGetUserByTokenParams{
			Token:         token,
			EncryptionKey: *s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, user, nil
	}
	if err != nil {
		return false, user, err
	}

	return true, user, nil
}
