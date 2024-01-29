package models

import (
	"time"
)

// Partner : database partner entity model
type Partner struct {
	UpdatedAt        time.Time
	CreatedAt        time.Time
	ID               string `gorm:"primary_key;size:255;not null"`
	Logo             string `gorm:"size:255"`
	SiteUrl          string `gorm:"size:255"`
	WhitelistDomains string `gorm:"size:2550"`
	Name             string `gorm:"size:255;not null"`
	ApiKey           string `gorm:"size:255;not null;unique;column:apiKey"`
	Active           int
}

// TableName sets the insert table name for this struct type
func (c *Partner) TableName() string {
	return "cdv_partners"
}
