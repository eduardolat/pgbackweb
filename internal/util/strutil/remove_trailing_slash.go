package strutil

// RemoveTrailingSlash removes the trailing slash from a string.
func RemoveTrailingSlash(str string) string {
	if len(str) > 0 && str[len(str)-1] == '/' {
		return str[:len(str)-1]
	}

	return str
}
