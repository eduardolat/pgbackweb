package pathutil

import "testing"

func TestBuildPath(t *testing.T) {
	t.Helper()

	tests := []struct {
		name       string
		prefix     string
		path       string
		expected   string
		shouldSkip bool // Skip if prefix was already set
	}{
		{
			name:     "no prefix configured",
			prefix:   "",
			path:     "/dashboard",
			expected: "/dashboard",
		},
		{
			name:     "with prefix - dashboard",
			prefix:   "/pgbackweb",
			path:     "/dashboard",
			expected: "/pgbackweb/dashboard",
		},
		{
			name:     "with prefix - api",
			prefix:   "/pgbackweb",
			path:     "/api/v1/health",
			expected: "/pgbackweb/api/v1/health",
		},
		{
			name:     "with prefix - root",
			prefix:   "/pgbackweb",
			path:     "",
			expected: "/pgbackweb",
		},
		{
			name:     "with prefix - auth",
			prefix:   "/pgbackweb",
			path:     "/auth/login",
			expected: "/pgbackweb/auth/login",
		},
		{
			name:     "empty prefix and empty path",
			prefix:   "",
			path:     "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset for each test
			pathPrefix = tt.prefix
			result := BuildPath(tt.path)
			if result != tt.expected {
				t.Errorf("BuildPath(%q) with prefix %q = %q, want %q", tt.path, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestGetPathPrefix(t *testing.T) {
	t.Helper()

	// Set a test prefix
	pathPrefix = "/test-prefix"

	result := GetPathPrefix()
	if result != "/test-prefix" {
		t.Errorf("GetPathPrefix() = %q, want %q", result, "/test-prefix")
	}

	// Reset
	pathPrefix = ""
}
