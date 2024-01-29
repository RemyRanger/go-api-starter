package models

import (
	"time"

	"gorm.io/gorm"
)

// Property : database property entity model
type Property struct {
	UpdateAt                    time.Time  `gorm:"column:update_at;type:datetime;"`
	IntegrationDate             time.Time  `gorm:"column:integration_date;type:datetime;"`
	YearOfConstruction          *string    `gorm:"column:year_of_construction;type:varchar;size:255;"`
	AvailabilityDate            *time.Time `gorm:"column:availability_date;type:datetime;"`
	PropertyTypeID              *int64     `gorm:"column:property_type_id;type:int;"`
	AgentID                     *int64     `gorm:"column:agent_id;type:int;"`
	DeveloperID                 *int64     `gorm:"column:developer_id;type:int;"`
	AgencyID                    *int64     `gorm:"column:agency_id;type:int;"`
	ProgramID                   *int64     `gorm:"column:program_id;type:int;"`
	Mobile                      *string    `gorm:"column:mobile;type:varchar;size:255;"`
	Email                       *string    `gorm:"column:email;type:varchar;size:255;"`
	Phone                       *string    `gorm:"column:phone;type:varchar;size:255;"`
	MobilePhone                 *string    `gorm:"column:mobile_phone;type:varchar;size:255;"`
	TypeCode                    *string    `gorm:"column:type_code;type:varchar;size:255;"`
	TypeLabel                   *string    `gorm:"column:type_label;type:varchar;size:255;"`
	Zip                         *string    `gorm:"column:zip;type:varchar;size:255;"`
	City                        *string    `gorm:"column:city;type:varchar;size:255;"`
	InseeCode                   *string    `gorm:"column:insee_code;type:varchar;size:255;"`
	ValidationStatus            *string    `gorm:"column:validation_status;type:varchar;size:255;"`
	GpsByLocationQuality        *string    `gorm:"column:gps_by_location_quality;type:varchar;size:3;"`
	Region                      *string    `gorm:"column:region;type:varchar;size:30;"`
	Department                  *string    `gorm:"column:department;type:varchar;size:30;"`
	Contact                     *string    `gorm:"column:contact;type:varchar;size:255;"`
	EntryDate                   *time.Time `gorm:"column:entry_date;type:datetime;"`
	Exposure                    *string    `gorm:"column:exposure;type:varchar;size:255;"`
	Kitchen                     *string    `gorm:"column:kitchen;type:varchar;size:255;"`
	Fibre                       *string    `gorm:"column:fibre;type:varchar;size:255;"`
	UrlPublicTariff             *string    `gorm:"column:url_public_tariff;type:varchar;size:255;"`
	Attic                       *string    `gorm:"column:attic;type:varchar;size:255;"`
	ModalityRentalCostRecovery  *string    `gorm:"column:modality_rental_cost_recovery;type:varchar;size:255;"`
	NearbyTransport             *string    `gorm:"column:nearby_transport;type:varchar;size:255;"`
	NearbySchool                *string    `gorm:"column:nearby_school;type:varchar;size:255;"`
	NearbyCommerce              *string    `gorm:"column:nearby_commerce;type:varchar;size:255;"`
	HeatingType                 *string    `gorm:"column:heating_type;type:varchar;size:255;"`
	District                    *string    `gorm:"column:district;type:varchar;size:255;"`
	DistrictOf                  *string    `gorm:"column:district_of;type:varchar;size:255;"`
	PayerFees                   *string    `gorm:"column:payer_fees;type:varchar;size:255;"`
	TenancyLaw                  *float64   `gorm:"column:tenancy_law;type:double;"`
	MainFeatures                *string    `gorm:"column:main_features;type:varchar;size:255;"`
	TypeStyle                   *string    `gorm:"column:type_style;type:varchar;size:255;"`
	Condominium                 *int64     `gorm:"column:condominium;type:tinyint;"`
	AdURL                       *string    `gorm:"column:ad_url;type:varchar;size:1024;"`
	Reference                   *string    `gorm:"column:reference;type:varchar;size:255;"`
	Street                      *string    `gorm:"column:street;type:varchar;size:255;"`
	Style                       *string    `gorm:"column:style;type:varchar;size:255;"`
	YearOfRenovation            *string    `gorm:"column:year_of_renovation;type:varchar;size:255;"`
	View                        *string    `gorm:"column:view;type:varchar;size:255;"`
	LotAvailability             *string    `gorm:"column:lot_availability;type:varchar;size:255;"`
	GeoZone                     *string    `gorm:"column:geo_zone;type:varchar;size:255;"`
	DeliveryYear                *string    `gorm:"column:delivery_year;type:varchar;size:255;"`
	Digicode                    *string    `gorm:"column:digicode;type:varchar;size:255;"`
	FlowID                      *string    `gorm:"column:flow_id;type:varchar;size:255;"`
	Environment                 *string    `gorm:"column:environment;type:varchar;size:255;"`
	HighSeasonComments          *string    `gorm:"column:high_season_comments;type:varchar;size:255;"`
	Outbuildings                *string    `gorm:"column:outbuildings;type:varchar;size:255;"`
	MiddleSeasonComments        *string    `gorm:"column:middle_season_comments;type:varchar;size:255;"`
	LowSeasonComments           *string    `gorm:"column:low_season_comments;type:varchar;size:255;"`
	Standing                    *string    `gorm:"column:standing;type:varchar;size:255;"`
	MandateRef                  *string    `gorm:"column:mandate_ref;type:varchar;size:255;"`
	MandateType                 *string    `gorm:"column:mandate_type;type:varchar;size:255;"`
	UrbanismCertificate         *int64     `gorm:"column:urbanism_certificate;type:tinyint;"`
	DpeLabelConsumption         *string    `gorm:"column:dpe_label_consumption;type:varchar;size:255;"`
	DpeLabelGaz                 *string    `gorm:"column:dpe_label_gaz;type:varchar;size:255;"`
	Heating                     *string    `gorm:"column:heating;type:varchar;size:255;"`
	HeatingEnergy               *string    `gorm:"column:heating_energy;type:varchar;size:255;"`
	HeatingMechanism            *string    `gorm:"column:heating_mechanism;type:varchar;size:255;"`
	InnerState                  *string    `gorm:"column:inner_state;type:varchar;size:255;"`
	AlurPercentageFeesIncluding *string    `gorm:"column:alur_percentage_fees_including;type:varchar;size:255;"`
	AlurUnionStatus             *string    `gorm:"column:alur_union_status;type:varchar;size:255;"`
	KitchenType                 *string    `gorm:"column:kitchen_type;type:varchar;size:255;"`
	FlowOrigin                  *string    `gorm:"column:flow_origin;type:varchar;size:50;"`
	MaxHeight                   *float64   `gorm:"column:max_height;type:double;"`
	VirtualVisit                *string    `gorm:"column:virtual_visit;type:varchar;size:1024;"`
	Description                 *string    `gorm:"column:description;type:text;size:4294967295;"`
	LotRef                      *string    `gorm:"column:lot_ref;type:varchar;size:255;"`
	Country                     *string    `gorm:"column:country;type:varchar;size:255;"`
	ParkingType                 *string    `gorm:"column:parking_type;type:varchar;size:255;"`
	OwnerType                   *string    `gorm:"column:owner_type;type:varchar;size:45;default:pro;"`
	Allotment                   *int64     `gorm:"column:allotment;type:tinyint;"`
	NbLots                      *int64     `gorm:"column:nb_lots;type:int;"`
	NbRooms                     *int64     `gorm:"column:nb_rooms;type:int;"`
	StayArea                    *int64     `gorm:"column:stay_area;type:int;"`
	NbBedRooms                  *int64     `gorm:"column:nb_bed_rooms;type:int;"`
	Depth                       *float64   `gorm:"column:depth;type:double;"`
	OccupiedLivingArea          *int64     `gorm:"column:occupied_living_area;type:int;"`
	FreeCommerceArea            *int64     `gorm:"column:free_commerce_area;type:int;"`
	OccupiedCommerceArea        *int64     `gorm:"column:occupied_commerce_area;type:int;"`
	Facade                      *float64   `gorm:"column:facade;type:double;"`
	Hnsa                        *int64     `gorm:"column:hnsa;type:int;"`
	Terrace                     *int64     `gorm:"column:terrace;type:tinyint;"`
	AlurNbLots                  *int64     `gorm:"column:alur_nb_lots;type:int;"`
	Far                         *int64     `gorm:"column:far;type:int;"`
	AlurCondominiumBackupPlan   *int64     `gorm:"column:alur_condominium_backup_plan;type:tinyint;"`
	FurniturePriceTtc           *int64     `gorm:"column:furniture_price_ttc;type:int;"`
	PrivateGarden               *int64     `gorm:"column:private_garden;type:tinyint;"`
	DpeValueGaz                 *float64   `gorm:"column:dpe_value_gaz;type:double;"`
	Balcony                     *int64     `gorm:"column:balcony;type:tinyint;"`
	DpeValueConsumption         *float64   `gorm:"column:dpe_value_consumption;type:double;"`
	LeadCertificate             *int64     `gorm:"column:lead_certificate;type:tinyint;"`
	DiagnosisAsbestos           *int64     `gorm:"column:diagnosis_asbestos;type:tinyint;"`
	DpeSubmitted                *int64     `gorm:"column:dpe_submitted;type:tinyint;"`
	DpeVirgin                   *int64     `gorm:"column:dpe_virgin;type:tinyint;"`
	Loggia                      *int64     `gorm:"column:loggia;type:tinyint;"`
	GreenSpaces                 *int64     `gorm:"column:green_spaces;type:tinyint;"`
	Intercom                    *int64     `gorm:"column:intercom;type:tinyint;"`
	AgencyFees                  *float64   `gorm:"column:agency_fees;type:double;"`
	Price                       *float64   `gorm:"column:price;type:double;"`
	Rent                        *float64   `gorm:"column:rent;type:double;"`
	Charges                     *float64   `gorm:"column:charges;type:double;"`
	GuaranteeDeposit            *float64   `gorm:"column:guarantee_deposit;type:double;"`
	LowSeasonPrice              *float64   `gorm:"column:low_season_price;type:double;"`
	MiddleSeasonPrice           *float64   `gorm:"column:middle_season_price;type:double;"`
	HighSeasonPrice             *float64   `gorm:"column:high_season_price;type:double;"`
	Floor                       *int64     `gorm:"column:floor;type:int;"`
	SwimmingPool                *int64     `gorm:"column:swimming_pool;type:tinyint;"`
	NbGarageSlots               *int64     `gorm:"column:nb_garage_slots;type:int;"`
	Garage                      *int64     `gorm:"column:garage;type:tinyint;"`
	LifeAnnuity                 *float64   `gorm:"column:life_annuity;type:double;"`
	Head1lifetimeAge            *int64     `gorm:"column:head1lifetime_age;type:int;"`
	Head2lifetimeAge            *int64     `gorm:"column:head2lifetime_age;type:int;"`
	FreeForSale                 *int64     `gorm:"column:free_for_sale;type:tinyint;"`
	Veranda                     *int64     `gorm:"column:veranda;type:tinyint;"`
	Basement                    *int64     `gorm:"column:basement;type:tinyint;"`
	TerraceArea                 *int64     `gorm:"column:terrace_area;type:int;"`
	BalconyArea                 *int64     `gorm:"column:balcony_area;type:int;"`
	Videocom                    *int64     `gorm:"column:videocom;type:tinyint;"`
	GardenArea                  *int64     `gorm:"column:garden_area;type:int;"`
	AreaMin                     *int64     `gorm:"column:area_min;type:int;"`
	AreaMax                     *int64     `gorm:"column:area_max;type:int;"`
	NbProperties                *int64     `gorm:"column:nb_properties;type:int;"`
	PriceMin                    *float64   `gorm:"column:price_min;type:double;"`
	NbWcs                       *int64     `gorm:"column:nb_wcs;type:int;"`
	NbGroundFloorRooms          *int64     `gorm:"column:nb_ground_floor_rooms;type:int;"`
	RentMin                     *float64   `gorm:"column:rent_min;type:double;"`
	NbBathrooms                 *int64     `gorm:"column:nb_bathrooms;type:int;"`
	Fireplace                   *int64     `gorm:"column:fireplace;type:tinyint;"`
	LivingArea                  *int64     `gorm:"column:living_area;type:int;"`
	NbAvailableProperties       *int64     `gorm:"column:nb_available_properties;type:int;"`
	NbRoomsMin                  *int64     `gorm:"column:nb_rooms_min;type:int;"`
	NbLevels                    *int64     `gorm:"column:nb_levels;type:int;"`
	Latitude                    *float64   `gorm:"column:latitude;type:decimal;"`
	Longitude                   *float64   `gorm:"column:longitude;type:decimal;"`
	TotalAreaOccupied           *int64     `gorm:"column:total_area_occupied;type:int;"`
	RentalYield                 *int64     `gorm:"column:rental_yield;type:int;"`
	Lot                         *int64     `gorm:"column:lot;type:tinyint;"`
	TypicalLot                  *int64     `gorm:"column:typical_lot;type:tinyint;"`
	FieldArea                   *int64     `gorm:"column:field_area;type:int;"`
	TotalFreeArea               *int64     `gorm:"column:total_free_area;type:int;"`
	NbParkings                  *int64     `gorm:"column:nb_parkings;type:int;"`
	Elevator                    *int64     `gorm:"column:elevator;type:tinyint;"`
	BatchLotsPossibleSale       *int64     `gorm:"column:batch_lots_possible_sale;type:tinyint;"`
	CondominiumExpenses         *float64   `gorm:"column:condominium_expenses;type:double;"`
	NbFloors                    *int64     `gorm:"column:nb_floors;type:int;"`
	Cellar                      *int64     `gorm:"column:cellar;type:tinyint;"`
	MonthlyRentIncCharge        *float64   `gorm:"column:monthly_rent_inc_charge;type:double;"`
	InventoryFees               *float64   `gorm:"column:inventory_fees;type:double;"`
	BuildingPermit              *int64     `gorm:"column:building_permit;type:tinyint;"`
	Furnished                   *int64     `gorm:"column:furnished;type:tinyint;"`
	PriceExclFees               *float64   `gorm:"column:price_excl_fees;type:double;"`
	SellerFees                  *int64     `gorm:"column:seller_fees;type:tinyint;"`
	BuyerFees                   *int64     `gorm:"column:buyer_fees;type:tinyint;"`
	PercentageSellerFees        *float64   `gorm:"column:percentage_seller_fees;type:double;"`
	PercentageBuyerFees         *float64   `gorm:"column:percentage_buyer_fees;type:double;"`
	AdditionnalRent             *float64   `gorm:"column:additionnal_rent;type:double;"`
	FreeLivingArea              *int64     `gorm:"column:free_living_area;type:int;"`
	NbWaterRooms                *int64     `gorm:"column:nb_water_rooms;type:int;"`
	RentMax                     *float64   `gorm:"column:rent_max;type:double;"`
	RenovationRate              *float64   `gorm:"column:renovation_rate;type:double;"`
	PropertyPriceTtc            *int64     `gorm:"column:property_price_ttc;type:int;"`
	PropertyPriceHt             *int64     `gorm:"column:property_price_ht;type:int;"`
	ProfitRate                  *float64   `gorm:"column:profit_rate;type:double;"`
	Connects                    *int64     `gorm:"column:connects;type:tinyint;"`
	FurniturePriceHt            *int64     `gorm:"column:furniture_price_ht;type:int;"`
	TaxRate                     *float64   `gorm:"column:tax_rate;type:double;"`
	Title                       string     `gorm:"column:title;type:text;size:4294967295;"`
	TypeParent                  string     `gorm:"column:type_parent;type:varchar;size:255;"`
	Service                     string     `gorm:"column:service;type:varchar;size:255;"`
	UbiflowID                   string     `gorm:"column:ubiflow_id;type:varchar;size:255;"`
	Developer                   Developer
	Agency                      Agency
	Plans                       []*Plan   `gorm:"foreignKey:PropertyID"`
	Medias                      []*Media  `gorm:"foreignKey:PropertyID"`
	Annexes                     []*Annexe `gorm:"foreignKey:PropertyID"`
	Program                     Program
	Area                        int64 `gorm:"column:area;type:int;"`
	Id                          int64 `gorm:"primary_key;column:id;type:int;"`
	Enabled                     int64 `gorm:"column:enabled;type:tinyint;"`
}

// TableName sets the insert table name for this struct type
func (cp *Property) TableName() string {
	return "cdv_property"
}

// BeforeUpdate : process just before update
func (cp *Property) BeforeUpdate(tx *gorm.DB) (err error) {
	cp.UpdateAt = time.Now()
	return
}

// BeforeCreate : process just before create
func (cp *Property) BeforeCreate(tx *gorm.DB) (err error) {
	cp.UpdateAt = time.Now()
	cp.IntegrationDate = time.Now()
	return
}
