package service

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/cron"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration"
	"github.com/eduardolat/pgbackweb/internal/service/auth"
	"github.com/eduardolat/pgbackweb/internal/service/backups"
	"github.com/eduardolat/pgbackweb/internal/service/databases"
	"github.com/eduardolat/pgbackweb/internal/service/destinations"
	"github.com/eduardolat/pgbackweb/internal/service/executions"
	"github.com/eduardolat/pgbackweb/internal/service/users"
)

type Service struct {
	AuthService         *auth.Service
	BackupsService      *backups.Service
	DatabasesService    *databases.Service
	DestinationsService *destinations.Service
	ExecutionsService   *executions.Service
	UsersService        *users.Service
}

func New(
	env *config.Env, dbgen *dbgen.Queries,
	cr *cron.Cron, ints *integration.Integration,
) *Service {
	authService := auth.New(env, dbgen)
	databasesService := databases.New(env, dbgen)
	destinationsService := destinations.New(env, dbgen)
	executionsService := executions.New(env, dbgen, ints)
	usersService := users.New(dbgen)
	backupsService := backups.New(dbgen, cr, executionsService)

	return &Service{
		AuthService:         authService,
		BackupsService:      backupsService,
		DatabasesService:    databasesService,
		DestinationsService: destinationsService,
		ExecutionsService:   executionsService,
		UsersService:        usersService,
	}
}
