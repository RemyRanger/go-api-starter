package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InvoiceAccount : database Project entity model
type InvoiceAccount struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string `gorm:"size:255;not null"`
	Apikeys       []Apikey
	Subscriptions []Subscription
	Company       Company
	CompanyUUID   uuid.UUID `gorm:"size:255;type:varchar"`
	UUID          uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (p *InvoiceAccount) TableName() string {
	return "invoice_accounts"
}

// BeforeCreate : This functions are called before creating Base
func (p *InvoiceAccount) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New()
	return
}

// CompanyID : CompanyID for copier mapping
func (p *InvoiceAccount) CompanyID(id string) {
	p.CompanyUUID, _ = uuid.Parse(id)
}
