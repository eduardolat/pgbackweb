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
	"github.com/eduardolat/pgbackweb/internal/service/oidc"
	"github.com/eduardolat/pgbackweb/internal/service/restorations"
	"github.com/eduardolat/pgbackweb/internal/service/users"
	"github.com/eduardolat/pgbackweb/internal/service/webhooks"
)

type Service struct {
	AuthService         *auth.Service
	BackupsService      *backups.Service
	DatabasesService    *databases.Service
	DestinationsService *destinations.Service
	ExecutionsService   *executions.Service
	OIDCService         *oidc.Service
	UsersService        *users.Service
	RestorationsService *restorations.Service
	WebhooksService     *webhooks.Service
}

// New constructs and initializes a Service instance with all component services.
// Returns the assembled Service or an error if OIDC service initialization fails.
func New(
	env config.Env, dbgen *dbgen.Queries,
	cr *cron.Cron, ints *integration.Integration,
) (*Service, error) {
	webhooksService := webhooks.New(dbgen)
	authService := auth.New(env, dbgen)
	oidcService, err := oidc.New(env, dbgen)
	if err != nil {
		return nil, err
	}
	databasesService := databases.New(env, dbgen, ints, webhooksService)
	destinationsService := destinations.New(env, dbgen, ints, webhooksService)
	executionsService := executions.New(env, dbgen, ints, webhooksService)
	usersService := users.New(dbgen)
	backupsService := backups.New(dbgen, cr, executionsService)
	restorationsService := restorations.New(
		dbgen, ints, executionsService, databasesService, destinationsService,
	)

	return &Service{
		AuthService:         authService,
		BackupsService:      backupsService,
		DatabasesService:    databasesService,
		DestinationsService: destinationsService,
		ExecutionsService:   executionsService,
		OIDCService:         oidcService,
		UsersService:        usersService,
		RestorationsService: restorationsService,
		WebhooksService:     webhooksService,
	}, nil
}
