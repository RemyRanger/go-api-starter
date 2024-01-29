package models

import (
	"time"
)

// DeveloperVisibility : database developerVisibility entity model
type DeveloperVisibility struct {
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;"`
	PartnerID      string    `gorm:"primary_key;column:partner_id;type:varchar;size:255;"`
	SubscriptionID string    `gorm:"column:subscription_id;type:varchar;size:255;"`
	Status         int64     `gorm:"column:status;type:tinyint;"`
	DeveloperID    int64     `gorm:"primary_key;column:developer_id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (c *DeveloperVisibility) TableName() string {
	return "cdv_developer_visibility"
}
