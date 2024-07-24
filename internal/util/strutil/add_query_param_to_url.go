package strutil

import (
	"fmt"
	nurl "net/url"
	"strings"
)

func AddQueryParamToUrl(url, key, value string) string {
	if url == "" {
		return ""
	}

	if key == "" || value == "" {
		return url
	}

	value = nurl.PathEscape(value)

	if !strings.Contains(url, "?") {
		return fmt.Sprintf("%s?%s=%s", url, key, value)
	}
	if strings.HasSuffix(url, "?") || strings.HasSuffix(url, "&") {
		return fmt.Sprintf("%s%s=%s", url, key, value)
	}
	return fmt.Sprintf("%s&%s=%s", url, key, value)
}
