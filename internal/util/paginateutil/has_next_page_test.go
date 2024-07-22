package paginateutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasNextPage(t *testing.T) {
	tests := []struct {
		name       string
		totalItems int
		limit      int
		offset     int
		want       bool
	}{
		{
			name:       "No items",
			totalItems: 0,
			limit:      10,
			offset:     0,
			want:       false,
		},
		{
			name:       "Exact number of items as limit",
			totalItems: 10,
			limit:      10,
			offset:     0,
			want:       false,
		},
		{
			name:       "More items than limit",
			totalItems: 15,
			limit:      10,
			offset:     0,
			want:       true,
		},
		{
			name:       "No more items beyond current page",
			totalItems: 20,
			limit:      10,
			offset:     10,
			want:       false,
		},
		{
			name:       "More items beyond current page",
			totalItems: 25,
			limit:      10,
			offset:     10,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasNextPage(tt.totalItems, tt.limit, tt.offset)
			assert.Equal(t, tt.want, got)
		})
	}
}
