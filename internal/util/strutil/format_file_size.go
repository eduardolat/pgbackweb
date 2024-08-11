package strutil

import "fmt"

// FormatFileSize pretty prints a file size (in bytes) to a human-readable format
//
// e.g. 1024 -> 1 KB
func FormatFileSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}

	if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	}

	if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
	}

	return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
}
