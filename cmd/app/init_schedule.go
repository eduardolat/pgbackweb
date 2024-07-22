package main

import (
	"github.com/eduardolat/pgbackweb/internal/cron"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/google/uuid"
)

func initSchedule(cr *cron.Cron, servs *service.Service) {
	/*
		Initial executions
	*/

	servs.ExecutionsService.SoftDeleteExpiredExecutions()
	servs.AuthService.DeleteOldSessions()

	/*
		Schedules
	*/

	err := cr.UpsertJob(uuid.New(), "UTC", "*/10 * * * *", func() {
		servs.ExecutionsService.SoftDeleteExpiredExecutions()
	})
	if err != nil {
		logger.FatalError(
			"error scheduling soft deletion of expired executions",
			logger.KV{"error": err},
		)
	}

	err = cr.UpsertJob(uuid.New(), "UTC", "*/10 * * * *", func() {
		servs.AuthService.DeleteOldSessions()
	})
	if err != nil {
		logger.FatalError(
			"error scheduling deletion of old sessions", logger.KV{"error": err},
		)
	}

	servs.BackupsService.ScheduleAll()
}
