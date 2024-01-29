package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SubscriptionExpanded : database subscription entity model
type SubscriptionExpanded struct {
	CreatedAt          time.Time
	UpdatedAt          time.Time
	StartDate          time.Time
	Status             string
	Description        string
	Type               string
	Product            Product
	InvoiceAccount     InvoiceAccount
	ActiveKeys         int64
	DurationDays       int32
	InvoiceAccountUUID uuid.UUID
	ProductUUID        uuid.UUID
	UUID               uuid.UUID
}

// Subscription : database subscription entity model
type Subscription struct {
	CreatedAt          time.Time
	UpdatedAt          time.Time
	StartDate          time.Time
	Description        string `gorm:"size:1024"`
	Type               string `gorm:"size:255;not null"`
	Status             string `gorm:"size:255;not null"`
	Product            Product
	InvoiceAccount     InvoiceAccount
	DurationDays       int32     `gorm:"not null"`
	InvoiceAccountUUID uuid.UUID `gorm:"size:255"`
	ProductUUID        uuid.UUID `gorm:"size:255"`
	UUID               uuid.UUID `gorm:"primary_key"`
}

// TableName sets the insert table name for this struct type
func (s *Subscription) TableName() string {
	return "subscriptions"
}

// BeforeCreate : This functions are called before creating Base
func (s *Subscription) BeforeCreate(tx *gorm.DB) (err error) {
	s.UUID = uuid.New()
	return
}

// InvoiceAccountID : InvoiceAccountID for copier mapping
func (s *Subscription) InvoiceAccountID(id string) {
	s.InvoiceAccountUUID, _ = uuid.Parse(id)
}

// ProductID : ProductID for copier mapping
func (s *Subscription) ProductID(id string) {
	s.ProductUUID, _ = uuid.Parse(id)
}
