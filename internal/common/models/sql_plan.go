package models

import (
	"database/sql"
	"time"
)

// Plan : database plan entity model
type Plan struct {
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;"`
	Position   sql.NullString `gorm:"column:position;type:varchar;size:255;"`
	Type       sql.NullString `gorm:"column:type;type:varchar;size:255;"`
	Folder     sql.NullString `gorm:"column:folder;type:varchar;size:255;"`
	Path       string         `gorm:"column:path;type:varchar;size:255;"`
	URL        sql.NullString `gorm:"column:url;type:varchar;size:255;"`
	FlowOrigin sql.NullString `gorm:"column:flow_origin;type:varchar;size:255;"`
	PropertyID sql.NullInt64  `gorm:"column:property_id;type:int;"`
	ProgramID  sql.NullInt64  `gorm:"column:program_id;type:int;"`
	Id         int64          `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (c *Plan) TableName() string {
	return "cdv_plan"
}
