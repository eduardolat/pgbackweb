package webhooks

import (
	"time"

	"github.com/google/uuid"
)

// DatabaseStatus represents the current health and metadata of a database.
type DatabaseStatus struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	PgVersion string    `json:"pg_version"`
	User      string    `json:"user"`
	Host      string    `json:"host"`
	Port      string    `json:"port"`
	DBName    string    `json:"dbname"`

	Healthy   bool      `json:"healthy"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`

	LastCheckedAt    *time.Time `json:"last_checked_at"`
	LastErrorMessage *string    `json:"last_error_message"`
	LastSuccess      *bool      `json:"last_success"`
}

// DestinationStatus represents the health and metadata of a destination (e.g., S3, backup target).
type DestinationStatus struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`

	Healthy   bool      `json:"healthy"`
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`

	LastCheckedAt    *time.Time `json:"last_checked_at"`
	LastErrorMessage *string    `json:"last_error_message"`
	LastSuccess      *bool      `json:"last_success"`

	Region     string `json:"region"`
	Endpoint   string `json:"endpoint"`
	BucketName string `json:"bucket_name"`
}

// ExecutionDetails holds information about a specific backup/test execution.
type ExecutionDetails struct {
	ID      uuid.UUID `json:"id"`
	Status  string    `json:"status"`
	Message string    `json:"message"`

	IsLocal bool `json:"is_local"`

	Path string `json:"path"`

	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`

	FileSize int64 `json:"file_size"`
}

// WebhookPayload aggregates all monitoring data sent via webhook.
type WebhookPayload struct {
	Database    DatabaseStatus    `json:"database"`
	Destination DestinationStatus `json:"destination"`
	Execution   ExecutionDetails  `json:"execution"`
	EventType   string            `json:"event_type"`
	Msg         string            `json:"msg"`
}
