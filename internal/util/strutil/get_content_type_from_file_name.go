package strutil

import "strings"

// GetContentTypeFromFileName returns the content type
// of a file based on its file extension from the name.
//
// If the file extension is not recognized, it returns
// "application/octet-stream".
func GetContentTypeFromFileName(fileName string) string {
	fileName = strings.ToLower(fileName)

	if strings.HasSuffix(fileName, ".pdf") {
		return "application/pdf"
	}

	if strings.HasSuffix(fileName, ".png") {
		return "image/png"
	}

	if strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".jpeg") {
		return "image/jpeg"
	}

	if strings.HasSuffix(fileName, ".gif") {
		return "image/gif"
	}

	if strings.HasSuffix(fileName, ".bmp") {
		return "image/bmp"
	}

	if strings.HasSuffix(fileName, ".json") {
		return "application/json"
	}

	if strings.HasSuffix(fileName, ".csv") {
		return "text/csv"
	}

	if strings.HasSuffix(fileName, ".xml") {
		return "application/xml"
	}

	if strings.HasSuffix(fileName, ".txt") {
		return "text/plain"
	}

	if strings.HasSuffix(fileName, ".html") {
		return "text/html"
	}

	if strings.HasSuffix(fileName, ".zip") {
		return "application/zip"
	}

	if strings.HasSuffix(fileName, ".sql") {
		return "application/sql"
	}

	return "application/octet-stream"
}
