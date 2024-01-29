package models

import (
	"time"
)

// Agency : database agency entity model
type Agency struct {
	UpdateAt        time.Time `gorm:"column:update_at;type:datetime;"`
	CompanyName     string    `gorm:"column:company_name;type:varchar;size:255;"`
	Siret           string    `gorm:"column:siret;type:varchar;size:255;"`
	Email           string    `gorm:"column:email;type:varchar;size:255;"`
	Phone           string    `gorm:"column:phone;type:varchar;size:255;"`
	Fax             string    `gorm:"column:fax;type:varchar;size:255;"`
	Street          string    `gorm:"column:street;type:varchar;size:255;"`
	Zip             string    `gorm:"column:zip;type:varchar;size:255;"`
	City            string    `gorm:"column:city;type:varchar;size:255;"`
	Country         string    `gorm:"column:country;type:varchar;size:255;"`
	Web             string    `gorm:"column:web;type:varchar;size:255;"`
	UbiflowID       string    `gorm:"column:ubiflow_id;type:varchar;size:255;"`
	PoiImage        string    `gorm:"column:poi_image;type:varchar;size:255;"`
	Logo            string    `gorm:"column:logo;type:varchar;size:255;"`
	Name            string    `gorm:"column:name;type:varchar;size:255;"`
	CostLink        string    `gorm:"column:cost_link;type:varchar;size:255;"`
	URLPublicTariff string    `gorm:"column:url_public_tariff;type:varchar;size:255;"`
	RemoteLogo      string    `gorm:"column:remote_logo;type:varchar;size:255;"`
	ID              int64     `gorm:"primary_key;column:id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (c *Agency) TableName() string {
	return "cdv_agency"
}
