package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Report : database Report entity model
type Report struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Name                string    `gorm:"size:124;not null"`
	GroupId             string    `gorm:"size:255;not null"`
	Status              string    `gorm:"size:20;not null"`
	Type                string    `gorm:"size:20;not null"`
	DemographyByPeriode string    `gorm:"size:4096;not null"`
	VentesByPeriode     string    `gorm:"size:4096;not null"`
	CreatedBy           string    `gorm:"size:124;not null"`
	Program             string    `gorm:"size:4096"`
	DeveloperPrograms   string    `gorm:"size:4096"`
	EcosystemPrograms   string    `gorm:"size:4096"`
	Comments            string    `gorm:"size:4096"`
	Prixm2ByPeriode     string    `gorm:"size:4096;not null"`
	Rentm2ByPeriode     string    `gorm:"size:4096;not null"`
	Prixm2ByType        string    `gorm:"size:4096;not null"`
	Longitude           float64   `gorm:"type:double;not null"`
	Latitude            float64   `gorm:"type:double;not null"`
	UUID                uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (f *Report) TableName() string {
	return "reports"
}

// BeforeCreate : This functions are called before creating Base
func (f *Report) BeforeCreate(tx *gorm.DB) (err error) {
	f.UUID = uuid.New()
	return
}
