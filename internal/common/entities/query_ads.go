package entities

import "time"

// Query represents an API query.
type QueryAds struct {
	NameQuery          *string
	IdQuery            *string
	ZipQuery           *string
	OrderBy            *string
	Flow               *string
	Service            *string
	TypeParent         *string
	MinIntegrationDate *time.Time
	MaxIntegrationDate *time.Time
	MinDeliveryDate    *time.Time
	MaxDeliveryDate    *time.Time
	SortDirDesc        *bool
	Enabled            *bool
	PageToken          *int
	PageSize           *int
}
