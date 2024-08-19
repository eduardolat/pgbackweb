package executions

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
)

type Service struct {
	env             *config.Env
	dbgen           *dbgen.Queries
	ints            *integration.Integration
	webhooksService *webhooks.Service
}

func New(
	env *config.Env, dbgen *dbgen.Queries, ints *integration.Integration,
	webhooksService *webhooks.Service,
) *Service {
	return &Service{
		env:             env,
		dbgen:           dbgen,
		ints:            ints,
		webhooksService: webhooksService,
	}
}
