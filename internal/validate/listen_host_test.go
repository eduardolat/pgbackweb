package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenHost(t *testing.T) {
	tests := []struct {
		name      string
		addr      string
		wantValid bool
	}{
		{
			name:      "valid ip address",
			addr:      "127.0.0.1",
			wantValid: true,
		},
		{
			name:      "valid ip address with CIDR",
			addr:      "192.168.1.1/24",
			wantValid: true,
		},
		{
			name:      "valid ip address zeros",
			addr:      "0.0.0.0",
			wantValid: true,
		},
		{
			name:      "invalid string",
			addr:      "invalid",
			wantValid: false,
		},
		{
			name:      "empty string",
			addr:      "",
			wantValid: false,
		},
		{
			name:      "invalid format with dots",
			addr:      "192.168.1",
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantValid, ListenHost(tt.addr))
		})
	}
}
