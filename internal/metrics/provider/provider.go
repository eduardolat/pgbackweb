package provider

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/logger"
	"github.com/eduardolat/pgbackweb/internal/metrics"
	"github.com/eduardolat/pgbackweb/internal/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitMetrics initializes the Prometheus metrics and starts the HTTP server if enabled.
func InitMetrics(env config.Env, db *sql.DB, servs *service.Service) {
	if !env.PBW_ENABLE_METRICS {
		return
	}

	// Register DB stats collector
	prometheus.MustRegister(collectors.NewDBStatsCollector(db, "pgbackweb"))

	// Periodically update health status
	go func() {
		for {
			ctx := context.Background()
			aggregatedDatabasesHealthy, aggregatedDestinationsHealthy := 1.0, 1.0
			healthyDatabasesCount, healthyDestinationsCount := 0.0, 0.0

			databases, err := servs.DatabasesService.GetAllDatabases(ctx)
			totalDatabasesCount := float64(len(databases))
			if err == nil {
				for _, db := range databases {
					status := 0.0
					if db.TestOk.Valid && db.TestOk.Bool {
						status = 1.0
						healthyDatabasesCount++
					} else {
						aggregatedDatabasesHealthy = 0.0
					}
					metrics.HealthStatus.WithLabelValues("database", db.Name).Set(status)
				}
			} else {
				aggregatedDatabasesHealthy = 0.0
				totalDatabasesCount = 0.0
			}

			destinations, err := servs.DestinationsService.GetAllDestinations(ctx)
			totalDestinationsCount := float64(len(destinations))
			if err == nil {
				for _, dest := range destinations {
					status := 0.0
					if dest.TestOk.Valid && dest.TestOk.Bool {
						status = 1.0
						healthyDestinationsCount++
					} else {
						aggregatedDestinationsHealthy = 0.0
					}
					metrics.HealthStatus.WithLabelValues("destination", dest.Name).Set(status)
				}
			} else {
				aggregatedDestinationsHealthy = 0.0
				totalDestinationsCount = 0.0
			}

			metrics.HealthStatus.WithLabelValues("databases", "all").Set(aggregatedDatabasesHealthy)
			metrics.HealthStatus.WithLabelValues("destinations", "all").Set(aggregatedDestinationsHealthy)

			// Update resource counts
			metrics.HealthyResourcesCount.WithLabelValues("database").Set(healthyDatabasesCount)
			metrics.HealthyResourcesCount.WithLabelValues("destination").Set(healthyDestinationsCount)
			metrics.TotalResourcesCount.WithLabelValues("database").Set(totalDatabasesCount)
			metrics.TotalResourcesCount.WithLabelValues("destination").Set(totalDestinationsCount)

			// Update backup task status
			allBackups, err := servs.BackupsService.GetAllBackups(ctx)
			activeBackupsCount, inactiveBackupsCount := 0.0, 0.0
			if err == nil {
				for _, backup := range allBackups {
					if backup.IsActive {
						activeBackupsCount++
					} else {
						inactiveBackupsCount++
					}
				}
			}
			metrics.BackupTasksStatus.WithLabelValues("active").Set(activeBackupsCount)
			metrics.BackupTasksStatus.WithLabelValues("inactive").Set(inactiveBackupsCount)

			time.Sleep(1 * time.Minute)
		}
	}()

	// Start standard Prometheus metrics server
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	address := fmt.Sprintf("%s:%s", env.PBW_METRICS_LISTEN_HOST, env.PBW_METRICS_LISTEN_PORT)
	go func() {
		logger.Info("metrics server started at http://" + address + "/metrics")
		if err := http.ListenAndServe(address, mux); err != nil {
			logger.Error("error starting metrics server", logger.KV{"error": err})
		}
	}()
}
