package validate

import (
	"strings"

	"github.com/adhocore/gronx"
)

// CronExpression validates a cron expression.
//
// It only supports 5 fields (minute, hour, day of month, month, day of week).
//
// It returns a boolean indicating whether the expression is valid or not.
func CronExpression(expression string) bool {
	fields := strings.Fields(expression)
	if len(fields) != 5 {
		return false
	}

	gron := gronx.New()
	return gron.IsValid(expression)
}
