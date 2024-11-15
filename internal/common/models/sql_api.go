package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Api struct {
	UpdatedAt   time.Time  `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	Description *string    `gorm:"size:1024"`
	ActiveOasID *uuid.UUID `gorm:"type:uuid"`
	DeletedAt   gorm.DeletedAt
	Title       string    `gorm:"size:256;not null"`
	Version     string    `gorm:"size:32;not null"`
	BaseUrl     string    `gorm:"size:1024;not null"`
	GatewayMode string    `gorm:"size:255;not null"`
	Status      string    `gorm:"size:255;not null"`
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (m *Api) TableName() string {
	return "sxc_api"
}
