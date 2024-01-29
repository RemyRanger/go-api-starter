package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow : database flow entity model
type Flow struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FlowOrigin  string    `gorm:"size:255;not null;unique"`
	JobName     string    `gorm:"size:255;not null;unique"`
	Status      string    `gorm:"size:255;not null"`
	Name        string    `gorm:"size:255;not null;unique"`
	Logo        string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:1024;not null"`
	TypeOffers  string    `gorm:"size:255;not null"`
	Format      string    `gorm:"size:255;not null"`
	UUID        uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (f *Flow) TableName() string {
	return "flows"
}

// BeforeCreate : This functions are called before creating Base
func (f *Flow) BeforeCreate(tx *gorm.DB) (err error) {
	f.UUID = uuid.New()
	return
}
