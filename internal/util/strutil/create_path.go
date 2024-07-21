package strutil

// CreatePath creates a path from the given parts.
//
// If addLeadingSlash is true, a leading slash will be added to the path.
//
// All colliding slashes will be converted to a single slash.
func CreatePath(addLeadingSlash bool, parts ...string) string {
	var path string

	for i, part := range parts {
		if part == "" {
			continue
		}

		cleanPart := RemoveLeadingSlash(part)
		path += "/" + cleanPart

		// Remove trailing slashes for all parts except the last one
		if i != len(parts)-1 {
			path = RemoveTrailingSlash(path)
		}
	}

	if !addLeadingSlash {
		path = RemoveLeadingSlash(path)
	}

	return path
}
