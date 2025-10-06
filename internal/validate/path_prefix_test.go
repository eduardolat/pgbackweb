package validate

import "testing"

func TestPathPrefix(t *testing.T) {
	t.Helper()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string is valid",
			input:    "",
			expected: true,
		},
		{
			name:     "valid simple path",
			input:    "/api",
			expected: true,
		},
		{
			name:     "valid complex path",
			input:    "/pgbackweb",
			expected: true,
		},
		{
			name:     "valid nested path",
			input:    "/app/v1",
			expected: true,
		},
		{
			name:     "valid deep nested path",
			input:    "/api/app/v1",
			expected: true,
		},
		{
			name:     "invalid - doesn't start with slash",
			input:    "api",
			expected: false,
		},
		{
			name:     "invalid - ends with slash",
			input:    "/api/",
			expected: false,
		},
		{
			name:     "invalid - only slash",
			input:    "/",
			expected: false,
		},
		{
			name:     "invalid - contains space",
			input:    "/api path",
			expected: false,
		},
		{
			name:     "invalid - contains tab",
			input:    "/api\tpath",
			expected: false,
		},
		{
			name:     "invalid - contains newline",
			input:    "/api\npath",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PathPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("PathPrefix(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
