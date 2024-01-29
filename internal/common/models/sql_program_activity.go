package models

import (
	"time"
)

// ProgramActivity : database program activity entity model
type ProgramActivity struct {
	UpdateAt time.Time
	Id       int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (p *ProgramActivity) TableName() string {
	return "cdv_program"
}
