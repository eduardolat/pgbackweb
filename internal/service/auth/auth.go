package auth

import (
	"time"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

const (
	maxSessionAge = time.Hour * 12
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
