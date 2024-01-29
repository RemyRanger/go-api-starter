package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ------------------- FORMAT MODELS

// EstimationResult defines model for EstimationResult.
type EstimationResult struct {
	Id                string  `json:"id"`
	ConfidenceIndice  int64   `json:"confidenceIndice"`
	EstimatedPrice    float64 `json:"estimatedPrice"`
	EstimatedPriceMax float64 `json:"estimatedPriceMax"`
	EstimatedPriceMin float64 `json:"estimatedPriceMin"`
}

// UUID : UUID for copier mapping
func (e *EstimationResult) UUID(uuid uuid.UUID) {
	e.Id = uuid.String()
}

// ------------------- DATABASE MODELS

// EstimationRequest : database EstimationRequest entity model
type EstimationRequest struct {
	CreatedAt    time.Time
	PropertyType string `gorm:"size:50;not null"`
	Estimation   Estimation
	Rooms        int64     `gorm:"not null"`
	Area         float64   `gorm:"not null"`
	Floor        int64     `gorm:"not null"`
	Latitude     float64   `gorm:"not null"`
	Longitude    float64   `gorm:"not null"`
	UUID         uuid.UUID `gorm:"primary_key"`
}

// BeforeCreate : This functions are called before creating Base
func (e *EstimationRequest) BeforeCreate(tx *gorm.DB) (err error) {
	e.UUID = uuid.New()
	return
}

// Estimation : database estimation entity model
type Estimation struct {
	CreatedAt             time.Time
	CreatedBy             string    `gorm:"size:50;not null"`
	EstimatedPriceMax     float64   `gorm:"not null"`
	EstimatedPrice        float64   `gorm:"not null"`
	EstimatedPriceMin     float64   `gorm:"not null"`
	ConfidenceIndice      int64     `gorm:"not null"`
	EstimationRequestUUID uuid.UUID `gorm:"not null;unique"`
	UUID                  uuid.UUID `gorm:"primary_key"`
}

// BeforeCreate : This functions are called before creating Base
func (e *Estimation) BeforeCreate(tx *gorm.DB) (err error) {
	e.UUID = uuid.New()
	return
}
