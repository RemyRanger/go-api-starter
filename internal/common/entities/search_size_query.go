package entities

// SearchSizeQuery defines model for search_size_query.
type SearchSizeQuery struct {
	TextQuery *string
	Size      *int64
}

// NewSearchSizeQuery : create new SearchYearsGeo entity
func NewSearchSizeQuery() *SearchSizeQuery {
	assignSize := int64(5)
	return &SearchSizeQuery{
		Size: &assignSize,
	}
}
