package entities

// SearchYearsGeo defines model for search_years_geo.
type SearchYearsZipCode struct {
	ZipCode string
	NbYears int64
}

// NewSearchYearsGeo : create new SearchYearsGeo entity
func NewSearchYearsZipCode() *SearchYearsZipCode {
	return &SearchYearsZipCode{
		NbYears: int64(5),
	}
}
