package entities

// SearchSuggestPrograms defines model for search_suggest_programs.
type SearchSuggestPrograms struct {
	Geometries struct {
		BoundingBox *BoundingBox
		Circle      *Circle
		Polygon     *Polygon
	}
	ConstructionType *string
	Page             *int64
	PageLimit        *int64
	AreaMax          int64
	AreaMin          int64
	Rooms            int64
}
