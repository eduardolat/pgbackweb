package strutil

// RemoveLeadingSlash removes the leading slash from a string.
func RemoveLeadingSlash(str string) string {
	if len(str) > 0 && str[0] == '/' {
		return str[1:]
	}

	return str
}
