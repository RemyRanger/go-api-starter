package models

import (
	"time"
)

// Defines values for ScpiType.
const (
	Pending ScpiVisibilityStatus = 0

	Validated ScpiVisibilityStatus = 1

	Deleted ScpiVisibilityStatus = 2
)

type ScpiVisibilityStatus int

// ScpiVisibility : database scpiVisibility entity model
type ScpiVisibility struct {
	ValidatedAt    time.Time            `gorm:"column:validated_at;type:datetime;"`
	CreatedAt      time.Time            `gorm:"column:created_at;type:datetime;"`
	UpdatedAt      time.Time            `gorm:"column:updated_at;type:datetime;"`
	PartnerID      string               `gorm:"primary_key;column:partner_id;type:varchar;size:255;"`
	SubscriptionID string               `gorm:"column:subscription_id;type:varchar;size:255;"`
	ManagerID      int64                `gorm:"column:manager_id;type:int;"`
	Status         ScpiVisibilityStatus `gorm:"column:status;type:int;"`
	Enabled        int64                `gorm:"column:enabled;type:tinyint;"`
	Shared         int64                `gorm:"column:shared;type:tinyint;"`
	ScpiID         int64                `gorm:"primary_key;column:scpi_id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (c *ScpiVisibility) TableName() string {
	return "cdv_scpi_visibility"
}
