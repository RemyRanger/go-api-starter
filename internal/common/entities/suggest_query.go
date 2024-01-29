package entities

// SuggestQuery defines model for suggestQuery.
type SuggestQuery struct {
	AreaMax         *int64
	AreaMin         *int64
	NbRoomsMax      *int64
	NbRoomsMin      *int64
	PageSize        *int64
	PriceMax        *int64
	PriceMin        *int64
	SameDeveloperId *bool
	SameZip         *bool
}

// Bind : just a post-process after a decode..
func NewSuggestQuery() *SuggestQuery {
	assignPageSize := int64(9)
	assignAreaMin := int64(9)
	return &SuggestQuery{
		PageSize: &assignPageSize,
		AreaMin:  &assignAreaMin,
	}
}
