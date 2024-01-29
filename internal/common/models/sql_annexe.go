package models

// Annexe : database Annexe entity model
type Annexe struct {
	Availability *string `gorm:"size:255"`
	Description  *string `gorm:"size:2048"`
	Area         *int64
	Price        *int64 `gorm:"column:price_ttc;type:int;"`
	PriceHt      *int64 `gorm:"column:price_ht;type:int;"`
	Type         string `gorm:"size:255"`
	Id           string `gorm:"primary_key;type:varchar"`
	Reference    string `gorm:"column:flow_id;type:varchar;"`
	PropertyID   int64  `gorm:"column:property_id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (a *Annexe) TableName() string {
	return "cdv_annexe"
}
