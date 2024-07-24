package databases

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration"
)

type Service struct {
	env   *config.Env
	dbgen *dbgen.Queries
	ints  *integration.Integration
}

func New(
	env *config.Env, dbgen *dbgen.Queries, ints *integration.Integration,
) *Service {
	return &Service{
		env:   env,
		dbgen: dbgen,
		ints:  ints,
	}
}
