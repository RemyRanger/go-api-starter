package entities

// ListPriceMeterType defines model for listPriceMeterType.
type ListPriceMeterType []struct {
	Data *[]struct {
		Typology *PricemeterTypologie `json:"typology,omitempty"`
		Period   *string              `json:"period,omitempty"`
	} `json:"data,omitempty"`
	InseeCode        *string `json:"insee_code,omitempty"`
	PeriodLastUpdate *string `json:"period_last_update,omitempty"`
}

// PricemeterType defines model for pricemeterType.
type PricemeterType struct {
	PriceMeterNewLodgmentWithParking    *int64 `json:"price_meter_new_lodgment_with_parking,omitempty"`
	PriceMeterNewLodgmentWithoutParking *int64 `json:"price_meter_new_lodgment_without_parking,omitempty"`
	PriceMeterOldLodgment               *int64 `json:"price_meter_old_lodgment,omitempty"`
}

// PricemeterTypologie defines model for pricemeterTypologie.
type PricemeterTypologie struct {
	T1 *PricemeterType `json:"t1,omitempty"`
	T2 *PricemeterType `json:"t2,omitempty"`
	T3 *PricemeterType `json:"t3,omitempty"`
	T4 *PricemeterType `json:"t4,omitempty"`
	T5 *PricemeterType `json:"t5,omitempty"`
}
