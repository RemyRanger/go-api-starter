package entities

import "time"

// SearchProperties defines model for searchProperties.
type SearchProperties struct {
	PageSize        *int64
	PageToken       *int64
	OrderBy         *string
	SortDirDesc     *bool
	FilterEnable    *bool
	DeveloperId     *string
	Service         *string
	OwnerType       *string
	TypeParents     *[]string
	FlowOrigin      *string
	Age             *string
	NbRooms         *[]int64
	MinPrice        *int64
	MaxPrice        *int64
	MinArea         *int64
	MaxArea         *int64
	MinDeliveryDate *time.Time
	MaxDeliveryDate *time.Time
	MinCreatedDate  *time.Time
	MaxCreatedDate  *time.Time
	Labels          *[]string
	TaxesExemption  *[]string
	GeoDistances    *[]struct {
		Distance *float64
		Lat      *float64
		Lng      *float64
	}
	GeoBoundingBox *[]struct {
		Lat *float64
		Lng *float64
	}
}

// NewSearchProperties : create new SearchProperties entity
func NewSearchProperties() *SearchProperties {
	assignPageToken := int64(1)
	assignPageSize := int64(10)
	return &SearchProperties{
		PageToken: &assignPageToken,
		PageSize:  &assignPageSize,
	}
}
