package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Status of the user
type UserStatus string

// Defines values for UserStatus.
const (
	UserStatusCREATED UserStatus = "CREATED"

	UserStatusDISABLED UserStatus = "DISABLED"

	UserStatusENABLED UserStatus = "ENABLED"
)

// User : database user entity model
type User struct {
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Role        string     `gorm:"size:255;not null"`
	Status      UserStatus `gorm:"size:255;not null"`
	ProviderID  string     `gorm:"size:255;not null;unique"`
	FirstName   string     `gorm:"size:255;not null"`
	LastName    string     `gorm:"size:255;not null"`
	Email       string     `gorm:"size:255;not null;unique"`
	Mobile      string     `gorm:"size:255;not null"`
	JobRole     string     `gorm:"size:255;not null"`
	Gender      string     `gorm:"size:255;not null"`
	Company     Company
	CompanyUUID uuid.UUID `gorm:"size:255;type:varchar"`
	UUID        uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate : This functions are called before creating Base
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Status = UserStatus("CREATED")
	u.UUID = uuid.New()
	return
}

// CompanyID : CompanyID for copier mapping
func (u *User) CompanyID(id string) {
	u.CompanyUUID, _ = uuid.Parse(id)
}
