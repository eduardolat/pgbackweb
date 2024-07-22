package paginateutil

// PaginateResponse is the response for a paginated request.
type PaginateResponse struct {
	TotalItems      int  `json:"total_items"`
	TotalPages      int  `json:"total_pages"`
	ItemsPerPage    int  `json:"items_per_page"`
	PreviousPage    int  `json:"previous_page"`
	HasPreviousPage bool `json:"has_previous_page"`
	CurrentPage     int  `json:"current_page"`
	NextPage        int  `json:"next_page"`
	HasNextPage     bool `json:"has_next_page"`
}

// CreatePaginateResponse creates a PaginateResponse from
// the given parameters.
func CreatePaginateResponse(
	paginateParams PaginateParams,
	totalItems int,
) PaginateResponse {
	limit := paginateParams.Limit

	totalPages := totalItems / limit
	if totalItems%limit != 0 {
		totalPages++
	}

	currentPage := paginateParams.Page
	previousPage := currentPage - 1
	nextPage := currentPage + 1

	if previousPage <= 0 {
		previousPage = 0
	}

	if totalPages < nextPage {
		nextPage = 0
	}

	hasPreviousPage := previousPage > 0
	hasNextPage := nextPage > 0

	return PaginateResponse{
		TotalItems:      totalItems,
		TotalPages:      totalPages,
		ItemsPerPage:    limit,
		PreviousPage:    previousPage,
		HasPreviousPage: hasPreviousPage,
		CurrentPage:     currentPage,
		NextPage:        nextPage,
		HasNextPage:     hasNextPage,
	}
}
