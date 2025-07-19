package users

import "github.com/eduardolat/pgbackweb/internal/database/dbgen"

type Service struct {
	dbgen *dbgen.Queries
}

// New creates and returns a new Service instance using the provided database queries handler.
func New(dbgen *dbgen.Queries) *Service {
	return &Service{
		dbgen: dbgen,
	}
}

// IsOIDCUser checks if a user is authenticated via OIDC
func (s *Service) IsOIDCUser(user dbgen.User) bool {
	return user.OidcProvider.Valid && user.OidcSubject.Valid
}
