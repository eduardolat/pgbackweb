# paginateutil

This package provides a utility for paginating data. It should be used in conjunction with database queries to paginate results.

Use these utilities to paginate results and, above all, to return a common structure in the different places where pagination is performed, maintaining consistency throughout the project.

## Usage

### 1. Define your queries:

- **PaginateCount:** This function should return the total number of records that match the query.
- **Paginate:** This function should return the paginated records.

### 2. Create the offset from the request parameters:

Use the `CreateOffsetFromParams` function to create the offset from the request parameters needed to paginate the results.

### 3. Define the signature of your pagination function:

The signature of your wrapper function should be:

```go
type PaginateXYZParams struct {
  Page      int
  Limit     int
  ABCFilter sql.NullString
}

func PaginateXYZ(
  ctx context.Context,
  params PaginateXYZParams,
) (
  paginateutil.PaginateResponse,
  []XYZ,
  error
)
```

### Example

Refer to `internal/service/backups/paginate_backups.go` for an example of how to use the `paginateutil` package.

## Notes

- **Default Values:** Ensure to set reasonable default values for pagination parameters (`page` and `limit`) if they are not provided.
- **Common Response Structure:** Use `CreatePaginateResponse` to generate a common response structure that includes information about pagination, such as the total number of items, the current page, and the number of items per page.
