package entities

// ListAdsLot defines model for listAdsLot.
type ListAdsLot struct {
	Lots       *[]AdsLot   `json:"lots,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// AdsLot defines model for adsLot.
type AdsLot struct {
	Area             *float64
	BrochureUrl      *string
	ConstructionType *string
	DeliveryDate     *string
	IdProgram        *string
	Location         *Location
	Price            *int64
	Rooms            *int64
}
