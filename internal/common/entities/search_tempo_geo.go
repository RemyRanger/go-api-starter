package entities

// SearchTempoGeo defines model for SearchTempoGeo.
type SearchTempoGeo struct {
	Frequency    *string
	NbOccurence  *int64
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

// NewSearchTempoGeo : create new SearchTempoGeo entity
func NewSearchTempoGeo() *SearchTempoGeo {
	assignFrequency := "year"
	assignNbOccurence := int64(5)
	return &SearchTempoGeo{
		Frequency:   &assignFrequency,
		NbOccurence: &assignNbOccurence,
	}
}
