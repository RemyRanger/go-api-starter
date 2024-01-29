package models

// Invisible : database Invisible entity model
type Invisible struct {
	PropertyID int64 `gorm:"primary_key;type:int;not null"`
	UserID     int64 `gorm:"primary_key;type:int;not null"`
}

// TableName sets the insert table name for this struct type
func (i *Invisible) TableName() string {
	return "cdv_delete"
}
