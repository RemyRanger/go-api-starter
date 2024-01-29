package entities

// ListProgram defines model for listProgram.
type ListProgram struct {
	Pagination *Pagination         `json:"pagination,omitempty"`
	Programs   []*SituationProgram `json:"programs,omitempty"`
}

// Pagination defines model for pagination.
type Pagination struct {
	Page     *int64
	PageSize *int64
	Total    *int64
}

// Program defines model for program.
type SituationProgram struct {
	Benefits           *ProgramBenefits
	BrochureUrl        *string
	ConfidenceIndex    *float64
	ConstructionType   *string
	DeliveryDate       *string
	Distance           *int64
	ExternId           *string
	FundingType        *string
	Id                 *string
	InvestorRate       *float64
	Location           *Location
	Name               *string
	Price              *ProgramPrice
	PromoterId         *string
	PromoterName       *string
	PromoterProgramId  *string
	ReducedVatZone     *bool
	Rhythm             *ProgramRhythm
	SaleDate           *string
	T1                 *Typo
	T2                 *Typo
	T3                 *Typo
	T4                 *Typo
	T5                 *Typo
	TotalAvailableLots *int64
	TotalLots          *int64
}

// ProgramBenefits defines model for program-benefits.
type ProgramBenefits struct {
	Certifications   *string
	GreenArea        *bool
	HeatingEnergy    *string
	HeatingMechanism *string
	HeatingType      *string
	Labels           *string
	NbFloors         *int64
}

// Location defines model for location.
type Location struct {
	City       *string
	InseeCode  *string
	Latitude   *string
	Longitude  *string
	SectorName *string
	Street     *string
	Zip        *string
}

// ProgramPrice defines model for program-price.
type ProgramPrice struct {
	FullVatPricePerMeterParkingExcepted           *int64
	FullVatPricePerMeterParkingIncluded           *int64
	FullVatPricePerMeterSubsidizedParkingExcepted *int64
	FullVatPricePerMeterSubsidizedParkingIncluded *int64
	ReducedVatPricePerMeterParkingExcepted        *int64
	ReducedVatPricePerMeterParkingIncluded        *int64
}

// ProgramRhythm defines model for program-rhythm.
type ProgramRhythm struct {
	Crusing   *Rhythm
	Global    *Rhythm
	Ninetieth *Rhythm
	Start     *Rhythm
}

// Rhythm defines model for rhythm.
type Rhythm struct {
	Duration     *float64
	FlowRate     *float64
	SalePerMonth *float64
}

// Typo defines model for typo.
type Typo struct {
	AllVatPrice            *TypoPrice
	FullVatPrice           *TypoPrice
	FullVatPriceSubsidized *TypoPrice
	LivingArea             *TypoArea
	ReducedVatPrice        *TypoPrice
	Volume                 *TypoVolume
}

// TypoPrice defines model for typo-price.
type TypoPrice struct {
	PriceParkingExcepted *TypoPriceMmm
	PriceParkingIncluded *TypoPriceMmm
}

// TypoPriceMmm defines model for typo-price-mmm.
type TypoPriceMmm struct {
	Avg *int64
	Max *int64
	Min *int64
}

// TypoArea defines model for typo-area.
type TypoArea struct {
	Avg *float64
	Max *float64
	Min *float64
}

// TypoVolume defines model for typo-volume.
type TypoVolume struct {
	AvailableOffer *int64
	InitialStock   *int64
	Sale           *int64
}
