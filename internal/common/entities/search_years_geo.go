package entities

// SearchYearsGeo defines model for search_years_geo.
type SearchYearsGeo struct {
	NbYears      *int64
	GeoDistances *[]struct {
		Distance *float64
		Lat      *float64
		Lng      *float64
	}
	GeoBoundingBox *[]struct {
		Lat *float64
		Lng *float64
	}
}

// NewSearchYearsGeo : create new SearchYearsGeo entity
func NewSearchYearsGeo() *SearchYearsGeo {
	assignNbYears := int64(5)
	return &SearchYearsGeo{
		NbYears: &assignNbYears,
	}
}
