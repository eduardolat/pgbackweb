package service

import (
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
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

func New(dbgen *dbgen.Queries) *Service {
	authService := auth.New(dbgen)
	backupsService := backups.New(dbgen)
	databasesService := databases.New(dbgen)
	destinationsService := destinations.New(dbgen)
	executionsService := executions.New(dbgen)
	usersService := users.New(dbgen)

	return &Service{
		AuthService:         authService,
		BackupsService:      backupsService,
		DatabasesService:    databasesService,
		DestinationsService: destinationsService,
		ExecutionsService:   executionsService,
		UsersService:        usersService,
	}
}
