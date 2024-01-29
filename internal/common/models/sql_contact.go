package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Contact : database contact entity model
type Contact struct {
	UpdatedAt   time.Time
	CreatedAt   time.Time
	ProjectUUID *uuid.UUID `gorm:"size:255;type:varchar"`
	FlowUUID    *uuid.UUID `gorm:"size:255;type:varchar"`
	FirstName   string     `gorm:"size:255;not null"`
	LastName    string     `gorm:"size:255;not null"`
	Email       string     `gorm:"size:255"`
	Mobile      string     `gorm:"size:255"`
	Job         string     `gorm:"size:255"`
	Gender      string     `gorm:"size:255;not null"`
	Project     Project
	UUID        uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (c *Contact) TableName() string {
	return "contacts"
}

// BeforeCreate : This functions are called before creating Base
func (c *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New()
	return
}

// ProjectID : ProjectID for copier mapping
func (c *Contact) ProjectID(id *string) {
	if id != nil {
		projectUUID, _ := uuid.Parse(*id)
		c.ProjectUUID = &projectUUID
	}
}

// ProjectID : ProjectID for copier mapping
func (c *Contact) FlowID(id *string) {
	if id != nil {
		flowUUID, _ := uuid.Parse(*id)
		c.FlowUUID = &flowUUID
	}
}
