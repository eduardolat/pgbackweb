package main

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/cron"
	"github.com/eduardolat/pgbackweb/internal/database"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service"
)

func main() {
	env := config.GetEnv()

	cr, err := cron.New()
	if err != nil {
		logger.FatalError("error initializing cron scheduler", logger.KV{"error": err})
	}
	cr.Start()
	defer cr.Shutdown()

	db := database.Connect(env)
	defer db.Close()
	dbgen := dbgen.New(db)

	ints := integration.New()
	servs := service.New(env, dbgen, cr, ints)

	initSchedule(cr, servs)
}
