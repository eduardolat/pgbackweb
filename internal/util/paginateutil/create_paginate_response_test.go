package paginateutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePaginateResponse(t *testing.T) {
	tests := []struct {
		name           string
		paginateParams PaginateParams
		totalItems     int
		expected       PaginateResponse
	}{
		{
			name: "Test with totalItems divisible by limit",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  1,
			},
			totalItems: 100,
			expected: PaginateResponse{
				TotalItems:      100,
				TotalPages:      10,
				ItemsPerPage:    10,
				PreviousPage:    0,
				HasPreviousPage: false,
				CurrentPage:     1,
				NextPage:        2,
				HasNextPage:     true,
			},
		},
		{
			name: "Test with totalItems not divisible by limit",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  1,
			},
			totalItems: 105,
			expected: PaginateResponse{
				TotalItems:      105,
				TotalPages:      11,
				ItemsPerPage:    10,
				PreviousPage:    0,
				HasPreviousPage: false,
				CurrentPage:     1,
				NextPage:        2,
				HasNextPage:     true,
			},
		},
		{
			name: "Test with offset",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  3,
			},
			totalItems: 100,
			expected: PaginateResponse{
				TotalItems:      100,
				TotalPages:      10,
				ItemsPerPage:    10,
				PreviousPage:    2,
				HasPreviousPage: true,
				CurrentPage:     3,
				NextPage:        4,
				HasNextPage:     true,
			},
		},
		{
			name: "Test with zero totalItems",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  1,
			},
			totalItems: 0,
			expected: PaginateResponse{
				TotalItems:      0,
				TotalPages:      0,
				ItemsPerPage:    10,
				PreviousPage:    0,
				HasPreviousPage: false,
				CurrentPage:     1,
				NextPage:        0,
				HasNextPage:     false,
			},
		},
		{
			name: "Test with totalItems less than limit",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  1,
			},
			totalItems: 5,
			expected: PaginateResponse{
				TotalItems:      5,
				TotalPages:      1,
				ItemsPerPage:    10,
				PreviousPage:    0,
				HasPreviousPage: false,
				CurrentPage:     1,
				NextPage:        0,
				HasNextPage:     false,
			},
		},
		{
			name: "Test with offset greater than totalItems",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  11,
			},
			totalItems: 50,
			expected: PaginateResponse{
				TotalItems:      50,
				TotalPages:      5,
				ItemsPerPage:    10,
				PreviousPage:    10,
				HasPreviousPage: true,
				CurrentPage:     11,
				NextPage:        0,
				HasNextPage:     false,
			},
		},
		{
			name: "Test with limit greater than totalItems",
			paginateParams: PaginateParams{
				Limit: 100,
				Page:  1,
			},
			totalItems: 50,
			expected: PaginateResponse{
				TotalItems:      50,
				TotalPages:      1,
				ItemsPerPage:    100,
				PreviousPage:    0,
				HasPreviousPage: false,
				CurrentPage:     1,
				NextPage:        0,
				HasNextPage:     false,
			},
		},
		{
			name: "Test if last page exists when it's incomplete",
			paginateParams: PaginateParams{
				Limit: 10,
				Page:  5,
			},
			totalItems: 52,
			expected: PaginateResponse{
				TotalItems:      52,
				TotalPages:      6,
				ItemsPerPage:    10,
				PreviousPage:    4,
				HasPreviousPage: true,
				CurrentPage:     5,
				NextPage:        6,
				HasNextPage:     true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CreatePaginateResponse(tt.paginateParams, tt.totalItems)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
