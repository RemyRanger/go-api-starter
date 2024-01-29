package entities

import "time"

// SearchPrograms defines model for searchPrograms.
type SearchPrograms struct {
	QuerySearch     *string
	PageSize        *int64
	PageToken       *int64
	OrderBy         *string
	SortDirDesc     *bool
	FilterEnable    *bool
	FilterNotLinked *bool
	DeveloperId     *string
	NbRooms         *[]int64
	MinPrice        *int64
	MaxPrice        *int64
	MinArea         *int64
	MaxArea         *int64
	MinDeliveryDate *time.Time
	MaxDeliveryDate *time.Time
	Labels          *[]string
	TaxesExemption  *[]string
	TypeParents     *[]string
	GeoBoundingBox  *[]struct {
		Lat *float64
		Lng *float64
	}
}

// Bind : just a post-process after a decode..
func NewSearchPrograms() *SearchPrograms {
	assignPageToken := int64(1)
	assignPageSize := int64(10)
	return &SearchPrograms{
		PageToken: &assignPageToken,
		PageSize:  &assignPageSize,
	}
}
