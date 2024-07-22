package paginateutil

// PaginateResponse is the response for a paginated request.
type PaginateResponse struct {
	TotalItems      int  `json:"totalItems"`
	TotalPages      int  `json:"totalPages"`
	ItemsPerPage    int  `json:"itemsPerPage"`
	PreviousPage    int  `json:"previousPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	CurrentPage     int  `json:"currentPage"`
	NextPage        int  `json:"nextPage"`
	HasNextPage     bool `json:"hasNextPage"`
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
