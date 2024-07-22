package paginateutil

// CreateOffsetFromParams creates an offset from the given
// PaginateParams.
func CreateOffsetFromParams(paginateParams PaginateParams) int {
	if paginateParams.Page <= 0 || paginateParams.Limit <= 0 {
		return 0
	}

	return (paginateParams.Page - 1) * paginateParams.Limit
}
