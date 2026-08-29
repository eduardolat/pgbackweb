package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// BackupsTotal is the total number of backups executed
	BackupsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pgbackweb_backups_total",
			Help: "Total number of backups executed, partitioned by status, backup name, and database name.",
		},
		[]string{"status", "backup_name", "database_name"},
	)

	// BackupsRunning is the number of backups currently running
	BackupsRunning = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_backups_running",
			Help: "Number of backups currently running per backup and database.",
		},
		[]string{"backup_name", "database_name"},
	)

	// LastBackupDuration is the duration of the last backup in seconds
	LastBackupDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_last_backup_duration_seconds",
			Help: "Duration of the last backup in seconds per backup and database.",
		},
		[]string{"backup_name", "database_name"},
	)

	// LastBackupStatus reports the status of the last backup (1 for success, 0 for failure)
	LastBackupStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_last_backup_status",
			Help: "Status of the last backup (1 = success, 0 = failure) per backup and database.",
		},
		[]string{"backup_name", "database_name"},
	)

	// HealthStatus reports the health status of databases and destinations (1 for healthy, 0 for unhealthy)
	HealthStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_health_status",
			Help: "Health status of databases and destinations (1 = healthy, 0 = unhealthy) per instance.",
		},
		[]string{"type", "name"},
	)

	// HealthyResourcesCount reports the count of healthy databases or destinations
	HealthyResourcesCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_healthy_resources_count",
			Help: "Count of healthy databases or destinations.",
		},
		[]string{"type"},
	)

	// TotalResourcesCount reports the total count of databases or destinations
	TotalResourcesCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_total_resources_count",
			Help: "Total count of databases or destinations.",
		},
		[]string{"type"},
	)

	// BackupTasksStatus reports the count of active and inactive backup tasks
	BackupTasksStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pgbackweb_backup_tasks_status",
			Help: "Count of backup tasks by status (active or inactive).",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(BackupsTotal)
	prometheus.MustRegister(BackupsRunning)
	prometheus.MustRegister(LastBackupDuration)
	prometheus.MustRegister(LastBackupStatus)
	prometheus.MustRegister(HealthStatus)
	prometheus.MustRegister(HealthyResourcesCount)
	prometheus.MustRegister(TotalResourcesCount)
	prometheus.MustRegister(BackupTasksStatus)
}
