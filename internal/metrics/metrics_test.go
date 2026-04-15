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
	// Simulate some metric updates
	BackupsRunning.WithLabelValues("my-backup", "my-db").Inc()
	BackupsTotal.WithLabelValues("success", "my-backup", "my-db").Inc()
	LastBackupDuration.WithLabelValues("my-backup", "my-db").Set(123.45)
	LastBackupStatus.WithLabelValues("my-backup", "my-db").Set(1.0)
	HealthStatus.WithLabelValues("database", "my-db").Set(1.0)
	HealthyResourcesCount.WithLabelValues("database").Set(5)
	TotalResourcesCount.WithLabelValues("database").Set(10)
	BackupTasksStatus.WithLabelValues("active").Set(3)

	req := httptest.NewRequest("GET", "/metrics", nil)
	rr := httptest.NewRecorder()
	handler := promhttp.Handler()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	body, _ := io.ReadAll(rr.Body)
	resp := string(body)

	assert.True(t, strings.Contains(resp, "pgbackweb_backups_running{backup_name=\"my-backup\",database_name=\"my-db\"} 1"))
	assert.True(t, strings.Contains(resp, "pgbackweb_backups_total{backup_name=\"my-backup\",database_name=\"my-db\",status=\"success\"} 1"))
	assert.True(t, strings.Contains(resp, "pgbackweb_last_backup_duration_seconds{backup_name=\"my-backup\",database_name=\"my-db\"} 123.45"))
	assert.True(t, strings.Contains(resp, "pgbackweb_last_backup_status{backup_name=\"my-backup\",database_name=\"my-db\"} 1"))
	assert.True(t, strings.Contains(resp, "pgbackweb_health_status{name=\"my-db\",type=\"database\"} 1"))
	assert.True(t, strings.Contains(resp, "pgbackweb_healthy_resources_count{type=\"database\"} 5"))
	assert.True(t, strings.Contains(resp, "pgbackweb_total_resources_count{type=\"database\"} 10"))
	assert.True(t, strings.Contains(resp, "pgbackweb_backup_tasks_status{status=\"active\"} 3"))

	// Cleanup
	BackupsRunning.WithLabelValues("my-backup", "my-db").Dec()
}
