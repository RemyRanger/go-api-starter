package models

// Favorite : database Favorite entity model
type Favorite struct {
	PropertyID int64 `gorm:"primary_key;type:int;not null"`
	UserID     int64 `gorm:"primary_key;type:int;not null"`
}

// TableName sets the insert table name for this struct type
func (f *Favorite) TableName() string {
	return "cdv_favorite"
}
