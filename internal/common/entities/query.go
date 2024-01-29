package entities

// Query represents an API query.
type Query struct {
	SortDirDesc *bool
	OrderBy     *string
	PageToken   *int
	PageSize    *int
	GroupId     string
	Query       string
	Filter      string
}
