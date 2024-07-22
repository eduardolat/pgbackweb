package paginateutil

// PaginateParams are the parameters for a paginated request.
type PaginateParams struct {
	Limit int `query:"limit" form:"limit" json:"limit"`
	Page  int `query:"page" form:"page" json:"page"`
}
