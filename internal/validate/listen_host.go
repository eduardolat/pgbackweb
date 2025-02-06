package validate

import (
	"regexp"
)

// ListenHost validates if addr is a valid host to listen on.
func ListenHost(addr string) bool {
	re := regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}($|/[0-9]{2})$`)
	return re.MatchString(addr)
}
