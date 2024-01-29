package models

import (
	"time"
)

// Developer : database developer entity model
type Developer struct {
	UpdateAt     time.Time `gorm:"column:update_at;type:datetime;"`
	FlowOrigin   string    `gorm:"column:flow_origin;type:varchar;size:255;"`
	Name         string    `gorm:"column:name;type:varchar;size:255;"`
	CompanyName  string    `gorm:"column:company_name;type:varchar;size:255;"`
	Email        string    `gorm:"column:email;type:varchar;size:255;"`
	Phone        string    `gorm:"column:phone;type:varchar;size:255;"`
	Fax          string    `gorm:"column:fax;type:varchar;size:255;"`
	Street       string    `gorm:"column:street;type:varchar;size:255;"`
	Zip          string    `gorm:"column:zip;type:varchar;size:255;"`
	City         string    `gorm:"column:city;type:varchar;size:255;"`
	Country      string    `gorm:"column:country;type:varchar;size:255;"`
	Siret        string    `gorm:"column:siret;type:varchar;size:255;"`
	Web          string    `gorm:"column:web;type:varchar;size:255;"`
	Logo         string    `gorm:"column:logo;type:varchar;size:255;"`
	CrnCity      string    `gorm:"column:crn_city;type:varchar;size:100;"`
	CrnNumber    string    `gorm:"column:crn_number;type:varchar;size:100;"`
	UbiflowID    string    `gorm:"column:ubiflow_id;type:varchar;size:255;"`
	RemoteLogo   string    `gorm:"column:remote_logo;type:varchar;size:255;"`
	FlowOriginID string    `gorm:"column:flow_origin_id;type:varchar;size:255;"`
	Id           int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (d *Developer) TableName() string {
	return "cdv_developer"
}
