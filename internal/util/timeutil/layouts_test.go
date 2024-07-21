package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLayoutRFC3339(t *testing.T) {
	table := []struct {
		timeStr  string
		expected time.Time
	}{
		{
			timeStr:  "2022-12-31T23:59:59Z",
			expected: time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			timeStr:  "2022-12-31T20:59:59+01:00",
			expected: time.Date(2022, 12, 31, 19, 59, 59, 0, time.UTC),
		},
		{
			timeStr:  "2022-12-31T20:59:59-01:00",
			expected: time.Date(2022, 12, 31, 21, 59, 59, 0, time.UTC),
		},
	}

	for _, tt := range table {
		parsedTime, err := time.Parse(LayoutRFC3339, tt.timeStr)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, parsedTime.UTC())
	}
}

func TestLayoutRFC3339Nano(t *testing.T) {
	table := []struct {
		timeStr  string
		expected time.Time
	}{
		{
			timeStr:  "2022-12-31T23:59:59.123456789Z",
			expected: time.Date(2022, 12, 31, 23, 59, 59, 123456789, time.UTC),
		},
		{
			timeStr:  "2022-12-31T20:59:59.123456789+01:00",
			expected: time.Date(2022, 12, 31, 19, 59, 59, 123456789, time.UTC),
		},
		{
			timeStr:  "2022-12-31T20:59:59.123456789-01:00",
			expected: time.Date(2022, 12, 31, 21, 59, 59, 123456789, time.UTC),
		},
	}

	for _, tt := range table {
		parsedTime, err := time.Parse(LayoutRFC3339Nano, tt.timeStr)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, parsedTime.UTC())
	}
}

func TestLayoutInputDateTimeLocal(t *testing.T) {
	timeStr := "2022-12-31T23:59"
	expectedTime := time.Date(2022, 12, 31, 23, 59, 0, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutInputDateTimeLocal, timeStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}

func TestLayoutInputDate(t *testing.T) {
	dateStr := "2022-12-31"
	expectedTime := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutInputDate, dateStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}

func TestLayoutSlashDDMMYYYY(t *testing.T) {
	dateStr := "31/12/2022"
	expectedTime := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutSlashDDMMYYYY, dateStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}

func TestLayoutSlashYYYYMMDD(t *testing.T) {
	dateStr := "2022/12/31"
	expectedTime := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutSlashYYYYMMDD, dateStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}

func TestLayoutDashDDMMYYYY(t *testing.T) {
	dateStr := "31-12-2022"
	expectedTime := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutDashDDMMYYYY, dateStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}

func TestLayoutDashYYYYMMDD(t *testing.T) {
	dateStr := "2022-12-31"
	expectedTime := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutDashYYYYMMDD, dateStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}

func TestLayoutYYYYMMDDHHMMSS(t *testing.T) {
	dateStr := "20221231-235959"
	expectedTime := time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC)

	parsedTime, err := time.Parse(LayoutYYYYMMDDHHMMSS, dateStr)
	assert.NoError(t, err)
	assert.Equal(t, expectedTime, parsedTime)
}
