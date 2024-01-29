package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product : database product entity model
type Product struct {
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DocumentationLink string    `gorm:"size:255"`
	Status            string    `gorm:"size:255;not null"`
	Name              string    `gorm:"size:255;not null"`
	Code              string    `gorm:"size:255;not null"`
	Description       string    `gorm:"size:1024;not null"`
	DetailsLink       string    `gorm:"size:255"`
	Logo              string    `gorm:"size:255;not null"`
	Version           string    `gorm:"size:255;not null"`
	URL               string    `gorm:"size:255;not null"`
	Banner            string    `gorm:"size:255;not null"`
	Category          string    `gorm:"size:255;not null"`
	UUID              uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (p *Product) TableName() string {
	return "products"
}

// BeforeCreate : This functions are called before creating Base
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New()
	return
}
