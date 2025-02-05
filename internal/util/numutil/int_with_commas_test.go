package numutil

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntWithCommas(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{
			name:     "zero",
			input:    0,
			expected: "0",
		},
		{
			name:     "single digit",
			input:    5,
			expected: "5",
		},
		{
			name:     "two digits",
			input:    42,
			expected: "42",
		},
		{
			name:     "three digits",
			input:    999,
			expected: "999",
		},
		{
			name:     "four digits",
			input:    1000,
			expected: "1,000",
		},
		{
			name:     "five digits",
			input:    12345,
			expected: "12,345",
		},
		{
			name:     "six digits",
			input:    123456,
			expected: "123,456",
		},
		{
			name:     "seven digits",
			input:    1234567,
			expected: "1,234,567",
		},
		{
			name:     "eight digits",
			input:    12345678,
			expected: "12,345,678",
		},
		{
			name:     "nine digits",
			input:    123456789,
			expected: "123,456,789",
		},
		{
			name:     "ten digits",
			input:    1234567890,
			expected: "1,234,567,890",
		},
		{
			name:     "large number",
			input:    1234567890123456789,
			expected: "1,234,567,890,123,456,789",
		},
		{
			name:     "negative small",
			input:    -42,
			expected: "-42",
		},
		{
			name:     "negative large",
			input:    -12345678,
			expected: "-12,345,678",
		},
		{
			name:     "edge case 999999",
			input:    999999,
			expected: "999,999",
		},
		{
			name:     "edge case -1000",
			input:    -1000,
			expected: "-1,000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IntWithCommas(tt.input))
			if tt.input < math.MaxInt {
				assert.Equal(t, tt.expected, IntWithCommas(int(tt.input)))
			}
			if tt.input < math.MaxInt32 {
				assert.Equal(t, tt.expected, IntWithCommas(int32(tt.input)))
			}
		})
	}
}
