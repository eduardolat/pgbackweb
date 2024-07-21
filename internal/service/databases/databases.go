package databases

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

type Service struct {
	env   *config.Env
	dbgen *dbgen.Queries
}

func New(env *config.Env, dbgen *dbgen.Queries) *Service {
	return &Service{
		env:   env,
		dbgen: dbgen,
	}
}
