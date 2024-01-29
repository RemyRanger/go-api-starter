package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FlowVisibility : database flowVisibility entity model
type FlowVisibility struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Flow             Flow
	Subscription     Subscription
	Enabled          int64     `gorm:"column:enabled;type:tinyint;"`
	FlowUUID         uuid.UUID `gorm:"column:flow_uuid;type:varchar;size:255;"`
	SubscriptionUUID uuid.UUID `gorm:"column:subscription_uuid;type:varchar;size:255;"`
	UUID             uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (f *FlowVisibility) TableName() string {
	return "flow_visibility"
}

// BeforeCreate : This functions are called before creating Base
func (f *FlowVisibility) BeforeCreate(tx *gorm.DB) (err error) {
	f.UUID = uuid.New()
	return
}
