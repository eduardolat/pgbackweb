package main

import (
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/cron"
	"github.com/eduardolat/pgbackweb/internal/database"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/eduardolat/pgbackweb/internal/integration"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/eduardolat/pgbackweb/internal/view"
	"github.com/labstack/echo/v4"
)

func main() {
	env := config.GetEnv()

	cr, err := cron.New()
	if err != nil {
		logger.FatalError("error initializing cron scheduler", logger.KV{"error": err})
	}
	cr.Start()
	defer func() {
		if err := cr.Shutdown(); err != nil {
			logger.Error("error shutting down cron scheduler", logger.KV{"error": err})
		}
	}()

	db := database.Connect(env)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("error closing database connection", logger.KV{"error": err})
		}
	}()
	dbgen := dbgen.New(db)

	ints := integration.New()
	servs := service.New(env, dbgen, cr, ints)
	initSchedule(cr, servs)

	app := echo.New()
	app.HideBanner = true
	app.HidePort = true
	view.MountRouter(app, servs)

	logger.Info("server started at http://localhost:8085")
	if err := app.Start(":8085"); err != nil {
		logger.FatalError("error starting server", logger.KV{"error": err})
	}
}
