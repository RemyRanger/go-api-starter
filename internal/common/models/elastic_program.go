package models

import "time"

// ProgramES : ProgramES struct
type ProgramES struct {
	AgendaDeliveryDate           time.Time `gorm:"column:agenda_delivery_date" json:"agenda_delivery_date"`
	UpdatedAt                    time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt                    time.Time `gorm:"column:created_at" json:"created_at"`
	AgendaStartLaunch            time.Time `gorm:"column:agenda_start_launch" json:"agenda_start_launch"`
	AgendaCommercialisationDate  time.Time `gorm:"column:agenda_commercialisation_date" json:"agenda_commercialisation_date"`
	UbiflowID                    string    `gorm:"column:ubiflow_id" json:"ubiflow_id"`
	Title                        string    `gorm:"column:title" json:"title"`
	Description                  string    `gorm:"column:description" json:"description"`
	DeveloperName                string    `gorm:"column:developer_name" json:"developer_name"`
	Location                     string    `gorm:"column:location" json:"location"`
	FlowOrigin                   string    `gorm:"column:flow_origin" json:"flow_origin"`
	SuspendedForSubscriptionsIDs string    `gorm:"column:suspended_for_subscriptions_ids" json:"suspended_for_subscriptions_ids"`
	VisibleForSubscriptionsIDs   string    `gorm:"column:visible_for_subscriptions_ids" json:"visible_for_subscriptions_ids"`
	VisibleForKeysIDs            string    `gorm:"column:visible_for_keys_ids" json:"visible_for_keys_ids"`
	LinkedForSubscriptionsIDs    string    `gorm:"column:linked_for_subscriptions_ids" json:"linked_for_subscriptions_ids"`
	Delivery                     *string   `gorm:"column:delivery" json:"delivery"`
	ImagePaths                   string    `gorm:"column:image_paths" json:"image_paths"`
	NbRoom                       string    `gorm:"column:nb_room" json:"nb_room"`
	AddressZip                   string    `gorm:"column:address_zip" json:"address_zip"`
	AddressStreet                string    `gorm:"column:address_street" json:"address_street"`
	AddressCountry               string    `gorm:"column:address_country" json:"address_country"`
	AddressCity                  string    `gorm:"column:address_city" json:"address_city"`
	FlowID                       string    `gorm:"column:flow_id" json:"flow_id"`
	Label                        string    `gorm:"column:label" json:"label"`
	TaxExemption                 string    `gorm:"column:tax_exemption" json:"tax_exemption"`
	TypeParent                   string    `gorm:"column:type_parent" json:"type_parent"`
	PriceMax                     float64   `gorm:"column:price_max" json:"price_max"`
	AreaMin                      int64     `gorm:"column:area_min" json:"area_min"`
	AreaMax                      int64     `gorm:"column:area_max" json:"area_max"`
	NbRoomsMax                   int64     `gorm:"column:nb_rooms_max" json:"nb_rooms_max"`
	NbRoomsMin                   int64     `gorm:"column:nb_rooms_min" json:"nb_rooms_min"`
	NbAvailableProperties        int64     `gorm:"column:nb_available_properties" json:"nb_available_properties"`
	NbProperties                 int64     `gorm:"column:nb_properties" json:"nb_properties"`
	PriceMin                     float64   `gorm:"column:price_min" json:"price_min"`
	LinkID                       int64     `gorm:"column:link_id" json:"link_id"`
	Longitude                    float64   `gorm:"column:longitude" json:"longitude"`
	Latitude                     float64   `gorm:"column:latitude" json:"latitude"`
	DeveloperID                  int64     `gorm:"column:developer_id" json:"developer_id"`
	ID                           int64     `gorm:"primary_key;column:id" json:"id"`
}
