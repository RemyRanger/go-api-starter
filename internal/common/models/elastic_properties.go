package models

import "time"

// PropertyES : PropertyES struct
type PropertyES struct {
	AddressStreet                string    `gorm:"column:address_street" json:"address_street"`
	SuspendedForSubscriptionsIDs string    `gorm:"column:suspended_for_subscriptions_ids" json:"suspended_for_subscriptions_ids"`
	VisibleForSubscriptionsIDs   string    `gorm:"column:visible_for_subscriptions_ids" json:"visible_for_subscriptions_ids"`
	UbiflowID                    string    `gorm:"column:ubiflow_id" json:"ubiflow_id"`
	DeveloperName                string    `gorm:"column:developer_name" json:"developer_name"`
	Location                     string    `gorm:"column:location" json:"location"`
	ProgramName                  string    `gorm:"column:program_name" json:"program_name"`
	Description                  string    `gorm:"column:description" json:"description"`
	Reference                    string    `gorm:"column:reference" json:"reference"`
	FlowID                       string    `gorm:"column:flow_id" json:"flow_id"`
	ProgramUbiflowID             string    `gorm:"column:program_ubiflow_id" json:"program_ubiflow_id"`
	FlowOrigin                   string    `gorm:"column:flow_origin" json:"flow_origin"`
	VisibleForKeysIDs            string    `gorm:"column:visible_for_keys_ids" json:"visible_for_keys_ids"`
	Availability                 string    `gorm:"column:availability" json:"availability"`
	CreatedAt                    string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                    time.Time `gorm:"column:updated_at" json:"updated_at"`
	Title                        string    `gorm:"column:title" json:"title"`
	ImagePaths                   string    `gorm:"column:image_paths" json:"image_paths"`
	AddressZip                   string    `gorm:"column:address_zip" json:"address_zip"`
	Delivery                     string    `gorm:"column:delivery" json:"delivery"`
	Service                      string    `gorm:"column:service" json:"service"`
	TaxExemption                 string    `gorm:"column:tax_exemption" json:"tax_exemption"`
	TypeParent                   string    `gorm:"column:type_parent" json:"type_parent"`
	OwnerType                    string    `gorm:"column:owner_type" json:"owner_type"`
	AddressCity                  string    `gorm:"column:address_city" json:"address_city"`
	AddressCountry               string    `gorm:"column:address_country" json:"address_country"`
	Rent                         float64   `gorm:"column:rent" json:"rent"`
	Price                        float64   `gorm:"column:price" json:"price"`
	Longitude                    float64   `gorm:"column:longitude" json:"longitude"`
	Latitude                     float64   `gorm:"column:latitude" json:"latitude"`
	NbRooms                      int64     `gorm:"column:nb_rooms" json:"nb_rooms"`
	Area                         int64     `gorm:"column:area" json:"area"`
	ProgramID                    int64     `gorm:"column:program_id" json:"program_id"`
	DeveloperID                  int64     `gorm:"column:developer_id" json:"developer_id"`
	ID                           int64     `gorm:"column:id;primary_key" json:"id"`
}
