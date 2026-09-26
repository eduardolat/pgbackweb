package webhooks

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"text/template"

	"github.com/eduardolat/pgbackweb/internal/util/strutil"

	"time"
)

// ParsedDatabaseInfo holds the extracted components of the PostgreSQL URL.
type ParsedDatabaseInfo struct {
	User   string
	Host   string
	Port   string
	DBName string
}

// ParsePostgresURL extracts components from a PostgreSQL connection string URI.
func ParsePostgresURL(connString string) (ParsedDatabaseInfo, error) {
	// Parse the URI string
	u, err := url.Parse(connString)
	if err != nil {
		return ParsedDatabaseInfo{}, fmt.Errorf("failed to parse connection string: %w", err)
	}

	if u.Scheme != "postgresql" && u.Scheme != "postgres" {
		return ParsedDatabaseInfo{}, fmt.Errorf("unsupported database scheme: %s", u.Scheme)
	}

	host := u.Hostname()
	port := u.Port()
	if port == "" {
		port = "5432"
	}
	user := u.User.Username()
	dbName := strings.TrimPrefix(u.Path, "/")

	return ParsedDatabaseInfo{
		User:   user,
		Host:   host,
		Port:   port,
		DBName: dbName,
	}, nil
}

func RenderWebhookBody(payload WebhookPayload, bodyTemplate string) (string, error) {
	if bodyTemplate == "" {
		bodyTemplate = "{}"
	}

	funcMap := template.FuncMap{

		// A useful function for formatting time directly in the template
		"formatTime":     func(t time.Time) string { return t.Format(time.RFC3339) },
		"formatFileSize": strutil.FormatFileSize,
	}

	tmpl, err := template.New("webhookBody").Funcs(funcMap).Parse(bodyTemplate)
	if err != nil {
		return "", fmt.Errorf("error parsing webhook body template: %w", err)
	}

	// Execute the template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, payload)
	if err != nil {
		return "", fmt.Errorf("error executing webhook body template: %w", err)
	}

	return buf.String(), nil
}

func buildMessage(event eventType, payload WebhookPayload) string {
	switch event {
	case EventTypeDatabaseHealthy:
		return fmt.Sprintf(
			"Database '%s' is healthy as of %s.",
			payload.Database.Name, payload.Database.Timestamp.Format(time.RFC3339),
		)

	case EventTypeDatabaseUnhealthy:
		return fmt.Sprintf(
			"Database '%s' is UNHEALTHY. Error: %s",
			payload.Database.Name, payload.Database.Error,
		)

	case EventTypeDestinationHealthy:
		return fmt.Sprintf(
			"Destination '%s' is healthy as of %s.",
			payload.Destination.Name, payload.Destination.Timestamp.Format(time.RFC3339),
		)

	case EventTypeDestinationUnhealthy:
		return fmt.Sprintf(
			"Destination '%s' is UNHEALTHY. Error: %s",
			payload.Destination.Name, payload.Destination.Error,
		)

	case EventTypeExecutionSuccess:
		storageLocation := "Local storage"
		if !payload.Execution.IsLocal {
			storageLocation = fmt.Sprintf(
				"S3 bucket '%s' (%s)",
				payload.Destination.BucketName, payload.Destination.Region,
			)
		}

		return fmt.Sprintf(
			"Backup completed successfully for database '%s'. Stored in %s. File: %s (%s).",
			payload.Database.Name,
			storageLocation,
			payload.Execution.Path,
			strutil.FormatFileSize(payload.Execution.FileSize),
		)

	case EventTypeExecutionFailed:
		return fmt.Sprintf(
			"Backup FAILED for database '%s'. Error: %s",
			payload.Database.Name, payload.Execution.Message,
		)
	default:
		return "Unknown event"
	}
}
