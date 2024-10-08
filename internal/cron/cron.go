package cron

import (
	"fmt"

	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

// Cron is a wrapper around the gocron.Scheduler with the specific
// configuration for the project.
type Cron struct {
	scheduler gocron.Scheduler
}

// New creates a new instance of the Cron struct.
func New() (*Cron, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	logger.Info("cron scheduler initialized")

	return &Cron{
		scheduler: scheduler,
	}, nil
}

// UpsertJob adds a new job to the scheduler and it deletes the job first if
// it already exists.
func (c *Cron) UpsertJob(
	id uuid.UUID, timeZone string, cronExpression string,
	function any, parameters ...any,
) error {
	if err := c.RemoveJob(id); err != nil {
		return err
	}

	exp := fmt.Sprintf("CRON_TZ=%s %s", timeZone, cronExpression)
	_, err := c.scheduler.NewJob(
		gocron.CronJob(exp, false),
		gocron.NewTask(function, parameters...),
		gocron.WithIdentifier(id),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)

	return err
}

// RemoveJob removes a job from the scheduler only if it exists.
func (c *Cron) RemoveJob(id uuid.UUID) error {
	jobs := c.scheduler.Jobs()
	for _, job := range jobs {
		if job.ID() == id {
			return c.scheduler.RemoveJob(id)
		}
	}

	return nil
}

// Start starts the scheduler.
func (c *Cron) Start() {
	c.scheduler.Start()
}

// Shutdown stops the scheduler.
func (c *Cron) Shutdown() error {
	return c.scheduler.Shutdown()
}
