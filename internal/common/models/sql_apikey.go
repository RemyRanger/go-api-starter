package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Apikey : database apikey entity model
type Apikey struct {
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ValidityDate       *time.Time `gorm:"type:datetime"`
	Name               string     `gorm:"size:255;not null"`
	Token              string     `gorm:"type:varchar"`
	WhiteListDomains   *string    `gorm:"type:varchar;size:255"`
	InvoiceAccount     InvoiceAccount
	Subscription       Subscription
	UUID               uuid.UUID `gorm:"primary_key;type:varchar"`
	InvoiceAccountUUID uuid.UUID `gorm:"size:255;type:varchar"`
	SubscriptionUUID   uuid.UUID `gorm:"size:255;type:varchar"`
	Enabled            bool      `gorm:"type:boolean;not null"`
}

// TableName sets the insert table name for this struct type
func (a *Apikey) TableName() string {
	return "apikeys"
}

// BeforeCreate : This functions are called before creating Base
func (a *Apikey) BeforeCreate(tx *gorm.DB) (err error) {
	a.Enabled = true
	return
}

// InvoiceAccountID : InvoiceAccountID for copier mapping
func (a *Apikey) InvoiceAccountID(id string) {
	a.InvoiceAccountUUID, _ = uuid.Parse(id)
}

// SubscriptionID : SubscriptionID for copier mapping
func (a *Apikey) SubscriptionID(id string) {
	a.SubscriptionUUID, _ = uuid.Parse(id)
}

// WhiteList : WhiteList for copier mapping
func (a *Apikey) WhiteList(whiteList *string) {
	if whiteList != nil {
		a.WhiteListDomains = whiteList
	}
}
