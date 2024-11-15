package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Oas struct {
	UpdatedAt      time.Time `gorm:"not null"`
	CreatedAt      time.Time `gorm:"not null"`
	DeletedAt      gorm.DeletedAt
	ApiVersion     string `gorm:"size:256;not null"`
	OpenapiVersion string `gorm:"size:1024;not null"`
	JsonSpec       JSON   `gorm:"type:jsonb;not null"`
	Api            Api
	ApiID          uuid.UUID `gorm:"type:uuid"`
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (m *Oas) TableName() string {
	return "sxc_oas"
}
