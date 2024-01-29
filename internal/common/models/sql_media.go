package models

import (
	"time"

	"gitlab.com/cdv-projects/go-apis/internal/common/config"
)

// Media : database media entity model
type Media struct {
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;"`
	ProgramID   *int64    `gorm:"column:program_id;type:int;"`
	AgencyID    *int64    `gorm:"column:agency_id;type:int;"`
	DeveloperID *int64    `gorm:"column:developer_id;type:int;"`
	PropertyID  *int64    `gorm:"column:property_id;type:int;"`
	Path        string    `gorm:"column:path;type:varchar;size:255;"`
	URL         string    `gorm:"column:url;type:varchar;size:255;"`
	Position    string    `gorm:"column:position;type:varchar;size:255;"`
	Type        string    `gorm:"column:type;type:varchar;size:255;"`
	Folder      string    `gorm:"column:folder;type:varchar;size:255;"`
	Id          int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (m *Media) TableName() string {
	return "cdv_media"
}

// FormatPath url
func (m *Media) FormatPath(ubiflowID string) *Media {
	m.Path = config.Conf.Srv.ImageBaseURL + "/images/properties/" + ubiflowID + "/" + m.Path
	return m
}
