package auth

import (
	"context"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
)

func (s *Service) LoginOIDC(
	ctx context.Context, userID uuid.UUID, ip, userAgent string,
) (dbgen.AuthServiceLoginCreateSessionRow, error) {
	session, err := s.dbgen.AuthServiceLoginCreateSession(
		ctx, dbgen.AuthServiceLoginCreateSessionParams{
			UserID:        userID,
			Ip:            ip,
			UserAgent:     userAgent,
			Token:         uuid.NewString(),
			EncryptionKey: s.env.PBW_ENCRYPTION_KEY,
		},
	)
	if err != nil {
		return dbgen.AuthServiceLoginCreateSessionRow{}, err
	}

	return session, nil
}
