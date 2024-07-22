package paginateutil

// HasNextPage returns true if there are more pages to show.
func HasNextPage(
	totalItems int,
	limit int,
	offset int,
) bool {
	if totalItems <= 0 {
		return false
	}

	return totalItems > offset+limit
}
