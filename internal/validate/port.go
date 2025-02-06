package validate

import (
	"regexp"
	"strconv"
)

// Port validates if port is a valid port number.
func Port(port string) bool {
	re := regexp.MustCompile(`^\d{1,5}$`)
	if !re.MatchString(port) {
		return false
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	if portInt < 1 || portInt > 65535 {
		return false
	}

	return true
}
