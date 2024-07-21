package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadTimezones(t *testing.T) {
	for _, tz := range Timezones {
		t.Run(tz.Label, func(t *testing.T) {
			_, err := time.LoadLocation(tz.TzCode)
			assert.NoError(t, err)
		})
	}
}
