package entities

// ListPriceMeterPeriod defines model for listPriceMeterPeriod.
type ListPriceMeterPeriod []struct {
	Data             *[]PricemeterPeriod `json:"data,omitempty"`
	InseeCode        *string             `json:"insee_code,omitempty"`
	PeriodLastUpdate *string             `json:"period_last_update,omitempty"`
}

// PricemeterPeriod defines model for pricemeterPeriod.
type PricemeterPeriod struct {
	PriceMeterNewLodgmentWithParking    *int64  `json:"price_meter_new_lodgment_with_parking,omitempty"`
	PriceMeterNewLodgmentWithoutParking *int64  `json:"price_meter_new_lodgment_without_parking,omitempty"`
	PriceMeterOldLodgment               *int64  `json:"price_meter_old_lodgment,omitempty"`
	Period                              *string `json:"period,omitempty"`
}
