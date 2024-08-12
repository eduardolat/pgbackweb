package strutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		size     int64
		expected string
	}{
		{size: 0, expected: "0 B"},
		{size: 1, expected: "1 B"},
		{size: 1023, expected: "1023 B"},
		{size: 1024, expected: "1.00 KB"},
		{size: 1024*1024 - 10, expected: "1023.99 KB"},
		{size: 1024 * 1024, expected: "1.00 MB"},
		{size: 1024*1024*1024 - 10_000, expected: "1023.99 MB"},
		{size: 1024 * 1024 * 1024, expected: "1.00 GB"},
		{size: 1024*1024*1024*1024 - 10_000_000, expected: "1023.99 GB"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, FormatFileSize(test.size))
		})
	}
}
