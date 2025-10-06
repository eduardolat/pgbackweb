package validate

import "strings"

// PathPrefix validates that a path prefix is correctly formatted.
//
// Valid path prefixes:
// - Empty string (no prefix)
// - Must start with /
// - Must NOT end with /
// - No whitespace allowed
//
// Examples:
// - "" -> true (no prefix)
// - "/api" -> true
// - "/pgbackweb" -> true
// - "/app/v1" -> true
// - "api" -> false (doesn't start with /)
// - "/api/" -> false (ends with /)
// - "/ api" -> false (contains whitespace)
func PathPrefix(pathPrefix string) bool {
	// Empty string is valid (no prefix)
	if pathPrefix == "" {
		return true
	}

	// Must start with /
	if !strings.HasPrefix(pathPrefix, "/") {
		return false
	}

	// Must NOT end with /
	if strings.HasSuffix(pathPrefix, "/") {
		return false
	}

	// No whitespace allowed
	if strings.ContainsAny(pathPrefix, " \t\n\r") {
		return false
	}

	return true
}
