package entities

// ListSales defines model for listSales.
type ListSales []struct {
	Data             *[]Sale `json:"data,omitempty"`
	InseeCode        *string `json:"insee_code,omitempty"`
	PeriodLastUpdate *string `json:"period_last_update,omitempty"`
	Pinel            *Pinel  `json:"pinel,omitempty"`
}

// Pinel defines model for pinel.
type Pinel struct {
	Rent   *float64 `json:"rent,omitempty"`
	Zoning *string  `json:"zoning,omitempty"`
}

// Sale defines model for sale.
type Sale struct {
	Investments *struct {
		InvestorRate  *float64 `json:"investor_rate,omitempty"`
		ResidentRate  *float64 `json:"resident_rate,omitempty"`
		TotalInvestor *int64   `json:"total_investor,omitempty"`
		TotalResident *int64   `json:"total_resident,omitempty"`
	} `json:"investments,omitempty"`

	NewLodgmentTrends *struct {
		AvailableOfferVariationRate *float64 `json:"available_offer_variation_rate,omitempty"`
		PricePerMeterAvailableOffer *int64   `json:"price_per_meter_available_offer,omitempty"`
		PricePerMeterSales          *int64   `json:"price_per_meter_sales,omitempty"`
		SalesVariationRate          *float64 `json:"sales_variation_rate,omitempty"`
	} `json:"new_lodgment_trends,omitempty"`

	Rhythm *struct {
		AverageRhythm       *int64   `json:"average_rhythm,omitempty"`
		RhythmVariationRate *float64 `json:"rhythm_variation_rate,omitempty"`
	} `json:"rhythm,omitempty"`

	Volumes *struct {
		AvailableOffer *int64 `json:"available_offer,omitempty"`
		Commercialized *int64 `json:"commercialized,omitempty"`
		Sold           *int64 `json:"sold,omitempty"`
		TotalPrograms  *int64 `json:"total_programs,omitempty"`
	} `json:"volumes,omitempty"`
	Period *string `json:"period,omitempty"`
}
