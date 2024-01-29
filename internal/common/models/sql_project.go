package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project : database project entity model
type Project struct {
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"size:1024;not null"`
	Version     string `gorm:"size:255"`
	URL         string `gorm:"size:255"`
	Contacts    []Contact
	Company     Company
	CompanyUUID uuid.UUID `gorm:"size:255;type:varchar"`
	UUID        uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (p *Project) TableName() string {
	return "projects"
}

// BeforeCreate : This functions are called before creating Base
func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New()
	return
}

// BeforeCreate : This functions are called before creating Base
func (p *Project) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Model(&Contact{}).Where("project_uuid = ?", p.UUID).Unscoped().Delete(&Contact{}).Error
}

// CompanyID : CompanyID for copier mapping
func (p *Project) CompanyID(id string) {
	p.CompanyUUID, _ = uuid.Parse(id)
}
