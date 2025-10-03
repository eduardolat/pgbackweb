package pathutil

import "sync"

var (
	pathPrefix     string
	pathPrefixOnce sync.Once
)

// SetPathPrefix sets the path prefix once. This should be called during
// application initialization with the value from the environment config.
func SetPathPrefix(prefix string) {
	pathPrefixOnce.Do(func() {
		pathPrefix = prefix
	})
}

// GetPathPrefix returns the configured path prefix.
func GetPathPrefix() string {
	return pathPrefix
}

// BuildPath constructs a full path by prepending the configured path prefix
// to the given path. If no prefix is configured, returns the path as-is.
//
// Examples:
//   - BuildPath("/dashboard") with prefix "/pgbackweb" -> "/pgbackweb/dashboard"
//   - BuildPath("/dashboard") with no prefix -> "/dashboard"
//   - BuildPath("") with prefix "/pgbackweb" -> "/pgbackweb"
func BuildPath(path string) string {
	if pathPrefix == "" {
		return path
	}
	return pathPrefix + path
}
