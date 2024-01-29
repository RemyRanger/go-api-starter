package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RentEvolution : database RentEvolution entity model
type RentEvolution struct {
	CreatedAt time.Time
	Data      string    `gorm:"size:4096;not null"`
	ZipCode   string    `gorm:"not null"`
	NbYears   int64     `gorm:"not null"`
	UUID      uuid.UUID `gorm:"primary_key;type:varchar"`
}

// BeforeCreate : This functions are called before creating Base
func (e *RentEvolution) BeforeCreate(tx *gorm.DB) (err error) {
	e.UUID = uuid.New()
	return
}

// ListRentMeterPeriod defines model for listRentMeterPeriod.
type ListRentMeterPeriod struct {
	Data    *[]RentMeterPeriod `json:"data,omitempty"`
	ZipCode *string            `json:"zip_code,omitempty"`
}

// RentMeterPeriod defines model for rentMeterPeriod.
type RentMeterPeriod struct {
	// Loyer m² moyen.
	RentMeter *float64 `json:"rent_meter,omitempty"`
	Year      *string  `json:"year,omitempty"`
}
