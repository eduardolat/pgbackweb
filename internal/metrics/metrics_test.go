package metrics

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestMetricsExposure(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T)
		cleanup  func(t *testing.T)
		expected []string
	}{
		{
			name: "backup_metrics",
			setup: func(t *testing.T) {
				BackupsRunning.WithLabelValues("backup-"+t.Name(), "db-"+t.Name()).Inc()
				BackupsTotal.WithLabelValues("success", "backup-"+t.Name(), "db-"+t.Name()).Inc()
				LastBackupDuration.WithLabelValues("backup-"+t.Name(), "db-"+t.Name()).Set(123.45)
				LastBackupStatus.WithLabelValues("backup-"+t.Name(), "db-"+t.Name()).Set(1.0)
				HealthStatus.WithLabelValues("database", "db-"+t.Name()).Set(1.0)
				HealthyResourcesCount.WithLabelValues("database").Set(5)
				TotalResourcesCount.WithLabelValues("database").Set(10)
				BackupTasksStatus.WithLabelValues("active").Set(3)
			},
			cleanup: func(t *testing.T) {
				BackupsRunning.WithLabelValues("backup-"+t.Name(), "db-"+t.Name()).Dec()
				BackupsTotal.DeleteLabelValues("success", "backup-"+t.Name(), "db-"+t.Name())
				LastBackupDuration.DeleteLabelValues("backup-"+t.Name(), "db-"+t.Name())
				LastBackupStatus.DeleteLabelValues("backup-"+t.Name(), "db-"+t.Name())
				HealthStatus.DeleteLabelValues("database", "db-"+t.Name())
				HealthyResourcesCount.DeleteLabelValues("database")
				TotalResourcesCount.DeleteLabelValues("database")
				BackupTasksStatus.DeleteLabelValues("active")
			},
			expected: []string{
				"pgbackweb_backups_running",
				"pgbackweb_backups_total",
				"pgbackweb_last_backup_duration_seconds",
				"pgbackweb_last_backup_status",
				"pgbackweb_health_status",
				"pgbackweb_healthy_resources_count",
				"pgbackweb_total_resources_count",
				"pgbackweb_backup_tasks_status",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			defer tc.cleanup(t)

			req := httptest.NewRequest("GET", "/metrics", nil)
			rr := httptest.NewRecorder()
			handler := promhttp.Handler()
			handler.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
			body, _ := io.ReadAll(rr.Body)
			resp := string(body)

			for _, expected := range tc.expected {
				assert.True(t, strings.Contains(resp, expected), "expected metric %s not found in response", expected)
			}
		})
	}
}
