package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Company : database company entity model
type Company struct {
	UpdatedAt          time.Time
	CreatedAt          time.Time
	ZipCode            string    `gorm:"size:255;not null"`
	Status             string    `gorm:"size:100;not null"`
	Address1           string    `gorm:"size:255;not null"`
	Address2           string    `gorm:"size:255;not null"`
	BusinessDomain     string    `gorm:"size:255;not null"`
	Code               string    `gorm:"size:255;not null"`
	ContactLegal       string    `gorm:"size:255;not null"`
	Country            string    `gorm:"size:255;not null"`
	Description        string    `gorm:"size:1024;not null"`
	Email              string    `gorm:"size:255;not null"`
	Logo               string    `gorm:"size:255;not null"`
	Name               string    `gorm:"size:255;not null"`
	RegistrationNumber string    `gorm:"size:255;not null"`
	URL                string    `gorm:"size:255;not null"`
	Vat                string    `gorm:"size:255;not null"`
	UUID               uuid.UUID `gorm:"primary_key;type:varchar"`
	Root               bool      `gorm:"not null"`
}

// TableName sets the insert table name for this struct type
func (c *Company) TableName() string {
	return "companies"
}

// BeforeCreate : This functions are called before creating Base
func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New()
	return
}
