package paginateutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOffsetFromParams(t *testing.T) {
	tests := []struct {
		name     string
		params   PaginateParams
		expected int
	}{
		{
			name: "Test 1: Page 1, Limit 10",
			params: PaginateParams{
				Page:  1,
				Limit: 10,
			},
			expected: 0,
		},
		{
			name: "Test 2: Page 2, Limit 10",
			params: PaginateParams{
				Page:  2,
				Limit: 10,
			},
			expected: 10,
		},
		{
			name: "Test 3: Page 3, Limit 20",
			params: PaginateParams{
				Page:  3,
				Limit: 20,
			},
			expected: 40,
		},
		{
			name: "Test 4: Page 0, Limit 10",
			params: PaginateParams{
				Page:  0,
				Limit: 10,
			},
			expected: 0,
		},
		{
			name: "Test 5: Page 5, Limit 0",
			params: PaginateParams{
				Page:  5,
				Limit: 0,
			},
			expected: 0,
		},
		{
			name: "Test 6: Page 0, Limit 0",
			params: PaginateParams{
				Page:  0,
				Limit: 0,
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CreateOffsetFromParams(tt.params)
			assert.Equal(t, tt.expected, result)
		})
	}
}
