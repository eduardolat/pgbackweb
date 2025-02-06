package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	tests := []struct {
		name      string
		port      string
		wantValid bool
	}{
		{
			name:      "valid port number",
			port:      "8080",
			wantValid: true,
		},
		{
			name:      "valid minimum port",
			port:      "1",
			wantValid: true,
		},
		{
			name:      "valid maximum port",
			port:      "65535",
			wantValid: true,
		},
		{
			name:      "invalid minimum port",
			port:      "0",
			wantValid: false,
		},
		{
			name:      "invalid maximum port",
			port:      "65536",
			wantValid: false,
		},
		{
			name:      "invalid port - letters",
			port:      "abc",
			wantValid: false,
		},
		{
			name:      "invalid port - empty",
			port:      "",
			wantValid: false,
		},
		{
			name:      "invalid port - special chars",
			port:      "8080!",
			wantValid: false,
		},
		{
			name:      "invalid port - decimal",
			port:      "8080.1",
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantValid, Port(tt.port))
		})
	}
}
