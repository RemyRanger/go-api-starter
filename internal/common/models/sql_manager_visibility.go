package models

import (
	"time"
)

// ManagerVisibility : database managerVisibility entity model
type ManagerVisibility struct {
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;"`
	PartnerID      string    `gorm:"primary_key;column:partner_id;type:varchar;size:255;"`
	SubscriptionID string    `gorm:"column:subscription_id;type:varchar;size:255;"`
	Enabled        int64     `gorm:"column:enabled;type:tinyint;"`
	Shared         int64     `gorm:"column:shared;type:tinyint;"`
	ManagerID      int64     `gorm:"primary_key;column:manager_id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (c *ManagerVisibility) TableName() string {
	return "cdv_manager_visibility"
}
