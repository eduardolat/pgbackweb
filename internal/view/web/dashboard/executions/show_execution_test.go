package executions

import (
	"bytes"
	"database/sql"
	"strings"
	"testing"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadExecutionDropdownButton(t *testing.T) {
	executionID := uuid.New()

	t.Run("renders link for successful execution with path", func(t *testing.T) {
		node := downloadExecutionDropdownButton(dbgen.ExecutionsServicePaginateExecutionsRow{
			ID:     executionID,
			Status: "success",
			Path:   sql.NullString{String: "backup.zip", Valid: true},
		})
		require.NotNil(t, node)

		var got bytes.Buffer
		require.NoError(t, node.Render(&got))

		html := got.String()
		assert.Contains(t, html, "Download backup")
		assert.Contains(t, html, "/dashboard/executions/"+executionID.String()+"/download")
		assert.Contains(t, html, `target="_blank"`)
		assert.Contains(t, html, `rel="noopener noreferrer"`)
	})

	t.Run("hides link when execution is not downloadable", func(t *testing.T) {
		tests := []dbgen.ExecutionsServicePaginateExecutionsRow{
			{ID: executionID, Status: "running", Path: sql.NullString{String: "backup.zip", Valid: true}},
			{ID: executionID, Status: "failed", Path: sql.NullString{String: "backup.zip", Valid: true}},
			{ID: executionID, Status: "success"},
		}

		for _, test := range tests {
			t.Run(strings.TrimSpace(test.Status), func(t *testing.T) {
				assert.Nil(t, downloadExecutionDropdownButton(test))
			})
		}
	})
}
