package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCronExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       bool
	}{
		{
			name:       "Valid 5-field expression",
			expression: "0 12 * * 1-5",
			want:       true,
		},
		{
			name:       "Invalid: too few fields",
			expression: "0 12 * *",
			want:       false,
		},
		{
			name:       "Invalid: too many fields",
			expression: "0 12 * * 1-5 2023",
			want:       false,
		},
		{
			name:       "Invalid: incorrect minute",
			expression: "60 12 * * 1-5",
			want:       false,
		},
		{
			name:       "Valid: complex expression",
			expression: "*/15 2,8-17 * * 1-5",
			want:       true,
		},
		{
			name:       "Invalid: incorrect day of week",
			expression: "0 12 * * 8",
			want:       false,
		},
		{
			name:       "Valid: using names",
			expression: "0 12 * JAN-DEC MON-FRI",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CronExpression(tt.expression)
			assert.Equal(t, tt.want, got, "CronExpression(%v)", tt.expression)
		})
	}
}
