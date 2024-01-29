package models

import (
	"time"
)

// ProgramVisibility : database programVisibility entity model
type ProgramVisibility struct {
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;"`
	SubscriptionID string    `gorm:"column:subscription_id;type:varchar;size:255;"`
	PartnerID      string    `gorm:"primary_key;column:partner_id;type:varchar;size:255;"`
	ProgramID      int64     `gorm:"primary_key;column:program_id;type:int;"`
	ProgramLinkID  int64     `gorm:"column:program_link_id;type:int;"`
	Status         int64     `gorm:"column:status;type:tinyint;"`
}

// TableName sets the insert table name for this struct type
func (c *ProgramVisibility) TableName() string {
	return "cdv_program_visibility"
}
