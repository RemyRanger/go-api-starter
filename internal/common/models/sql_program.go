package models

import (
	"time"

	"gorm.io/gorm"
)

// Program : database program entity model
type Program struct {
	IntegrationDate             time.Time  `gorm:"column:integration_date;type:datetime;"`
	UpdateAt                    time.Time  `gorm:"column:update_at;type:datetime;"`
	DeliveryDate                time.Time  `gorm:"column:delivery_date;type:datetime;"`
	TaxExempDemessineLaw        *int64     `gorm:"column:tax_exemp_demessine_law;type:tinyint;"`
	Zip                         *string    `gorm:"column:zip;type:varchar;size:255;"`
	OfferEndDate                *time.Time `gorm:"column:offer_end_date;type:datetime;"`
	CommercialisationDate       *time.Time `gorm:"column:commercialisation_date;type:datetime;"`
	OfferStartDate              *time.Time `gorm:"column:offer_start_date;type:datetime;"`
	ContactMobile               *string    `gorm:"column:contact_mobile;type:varchar;size:255;"`
	Phone                       *string    `gorm:"column:phone;type:varchar;size:255;"`
	MobilePhone                 *string    `gorm:"column:mobile_phone;type:varchar;size:255;"`
	Street                      *string    `gorm:"column:street;type:varchar;size:255;"`
	DeliveryDateSemester        *string    `gorm:"column:delivery_date_semester;type:varchar;size:255;"`
	City                        *string    `gorm:"column:city;type:varchar;size:255;"`
	Country                     *string    `gorm:"column:country;type:varchar;size:255;"`
	ContactName                 *string    `gorm:"column:contact_name;type:varchar;size:255;"`
	ContactPhone                *string    `gorm:"column:contact_phone;type:varchar;size:255;"`
	Email                       *string    `gorm:"column:email;type:varchar;size:255;"`
	ContactFax                  *string    `gorm:"column:contact_fax;type:varchar;size:255;"`
	ContactMail                 *string    `gorm:"column:contact_mail;type:varchar;size:255;"`
	ContactOpeningHours         *string    `gorm:"column:contact_opening_hours;type:varchar;size:255;"`
	ContactStreet               *string    `gorm:"column:contact_street;type:varchar;size:255;"`
	ContactZip                  *string    `gorm:"column:contact_zip;type:varchar;size:255;"`
	ContactCity                 *string    `gorm:"column:contact_city;type:varchar;size:255;"`
	ContactAddress              *string    `gorm:"column:contact_address;type:varchar;size:255;"`
	Contact                     *string    `gorm:"column:contact;type:varchar;size:255;"`
	Digicode                    *string    `gorm:"column:digicode;type:varchar;size:255;"`
	Slice1                      *int64     `gorm:"column:slice1;type:tinyint;"`
	BalconyArea                 *int64     `gorm:"column:balcony_area;type:int;"`
	FlowOrigin                  *string    `gorm:"column:flow_origin;type:varchar;size:50;"`
	HeatingType                 *string    `gorm:"column:heating_type;type:varchar;size:255;"`
	HeatingEnergy               *string    `gorm:"column:heating_energy;type:varchar;size:255;"`
	HeatingMechanism            *string    `gorm:"column:heating_mechanism;type:varchar;size:255;"`
	FlowID                      *string    `gorm:"column:flow_id;type:varchar;size:255;"`
	Certifications              *string    `gorm:"column:certifications;type:varchar;size:255;"`
	AdURL                       *string    `gorm:"column:ad_url;type:varchar;size:255;"`
	VirtualVisit                *string    `gorm:"column:virtual_visit;type:varchar;size:255;"`
	OperationRef                *string    `gorm:"column:operation_ref;type:varchar;size:255;"`
	DpeLabelConsumption         *string    `gorm:"column:dpe_label_consumption;type:varchar;size:255;"`
	DpeLabelGaz                 *string    `gorm:"column:dpe_label_gaz;type:varchar;size:255;"`
	Description                 *string    `gorm:"column:description;type:text;size:4294967295;"`
	Environment                 *string    `gorm:"column:environment;type:varchar;size:255;"`
	Department                  *string    `gorm:"column:department;type:varchar;size:255;"`
	DeliveryDateYear            *string    `gorm:"column:delivery_date_year;type:varchar;size:255;"`
	LowConsumLabel              *int64     `gorm:"column:low_consum_label;type:tinyint;"`
	BioConstruction             *int64     `gorm:"column:bio_construction;type:tinyint;"`
	DeliveryDateMonth           *string    `gorm:"column:delivery_date_month;type:varchar;size:255;"`
	DeliveryDateDay             *string    `gorm:"column:delivery_date_day;type:varchar;size:255;"`
	KitchenType                 *string    `gorm:"column:kitchen_type;type:varchar;size:255;"`
	TypeCode                    *string    `gorm:"column:type_code;type:varchar;size:255;"`
	UbiflowID                   *string    `gorm:"column:ubiflow_id;type:varchar;size:255;"`
	Title                       *string    `gorm:"column:title;type:text;size:4294967295;"`
	Reference                   *string    `gorm:"column:reference;type:varchar;size:255;"`
	OfferTitle                  *string    `gorm:"column:offer_title;type:varchar;size:255;"`
	OfferMention                *string    `gorm:"column:offer_mention;type:varchar;size:255;"`
	OfferType                   *string    `gorm:"column:offer_type;type:varchar;size:255;"`
	BackFlowAt                  *time.Time `gorm:"column:back_flow_at;type:datetime;"`
	OfferDescription            *string    `gorm:"column:offer_description;type:varchar;size:255;"`
	DeliveryDateTrimester       *string    `gorm:"column:delivery_date_trimester;type:varchar;size:255;"`
	TaxExempRecenteredRobienLaw *int64     `gorm:"column:tax_exemp_recentered_robien_law;type:tinyint;"`
	HqenvLabel                  *int64     `gorm:"column:hqenv_label;type:tinyint;"`
	TaxExempMarlauxLaw          *int64     `gorm:"column:tax_exemp_marlaux_law;type:tinyint;"`
	HighPerfLabel               *int64     `gorm:"column:high_perf_label;type:tinyint;"`
	TaxExempLmnpStatus          *int64     `gorm:"column:tax_exemp_lmnp_status;type:tinyint;"`
	TaxExempHmlaw               *int64     `gorm:"column:tax_exemp_hmlaw;type:tinyint;"`
	TaxExempDuflotLaw           *int64     `gorm:"column:tax_exemp_duflot_law;type:tinyint;"`
	TaxExempBareOwnership       *int64     `gorm:"column:tax_exemp_bare_ownership;type:tinyint;"`
	TaxExempServicesResidence   *int64     `gorm:"column:tax_exemp_services_residence;type:tinyint;"`
	TaxExempTourismResidence    *int64     `gorm:"column:tax_exemp_tourism_residence;type:tinyint;"`
	TaxExempMicroRegime         *int64     `gorm:"column:tax_exemp_micro_regime;type:tinyint;"`
	ScellierLaw                 *int64     `gorm:"column:scellier_law;type:tinyint;"`
	TaxExempLmpStatus           *int64     `gorm:"column:tax_exemp_lmp_status;type:tinyint;"`
	EffinerEnergyLabel          *int64     `gorm:"column:effiner_energy_label;type:tinyint;"`
	Balcony                     *int64     `gorm:"column:balcony;type:tinyint;"`
	EntryDate                   *time.Time `gorm:"column:entry_date;type:datetime;"`
	Floor                       *int64     `gorm:"column:floor;type:int;"`
	Elevator                    *int64     `gorm:"column:elevator;type:tinyint;"`
	HousingEnvLabel             *int64     `gorm:"column:housing_env_label;type:tinyint;"`
	QualitelLabel               *int64     `gorm:"column:qualitel_label;type:tinyint;"`
	PromotelecLabel             *int64     `gorm:"column:promotelec_label;type:tinyint;"`
	Iso9001label                *int64     `gorm:"column:iso9001label;type:tinyint;"`
	Iso9002label                *int64     `gorm:"column:iso9002label;type:tinyint;"`
	Iso9003label                *int64     `gorm:"column:iso9003label;type:tinyint;"`
	ImmediateDelivery           *int64     `gorm:"column:immediate_delivery;type:tinyint;"`
	LastOpportunity             *int64     `gorm:"column:last_opportunity;type:tinyint;"`
	WorkInProgress              *int64     `gorm:"column:work_in_progress;type:tinyint;"`
	CommercialLaunch            *int64     `gorm:"column:commercial_launch;type:tinyint;"`
	Preview                     *int64     `gorm:"column:preview;type:tinyint;"`
	CompatibleSocialLoan        *int64     `gorm:"column:compatible_social_loan;type:tinyint;"`
	CompatibleLzi               *int64     `gorm:"column:compatible_lzi;type:tinyint;"`
	TaxExempRobienLaw           *int64     `gorm:"column:tax_exemp_robien_law;type:tinyint;"`
	TaxExempBessonLaw           *int64     `gorm:"column:tax_exemp_besson_law;type:tinyint;"`
	Furnished                   *int64     `gorm:"column:furnished;type:tinyint;"`
	NbAvailableProperties       *int64     `gorm:"column:nb_available_properties;type:int;"`
	LivingArea                  *int64     `gorm:"column:living_area;type:int;"`
	DeveloperID                 *int64     `gorm:"column:developer_id;type:int;"`
	RentMax                     *float64   `gorm:"column:rent_max;type:double;"`
	RentMin                     *float64   `gorm:"column:rent_min;type:double;"`
	NbProperties                *int64     `gorm:"column:nb_properties;type:int;"`
	Latitude                    *float64   `gorm:"column:latitude;type:double;"`
	Longitude                   *float64   `gorm:"column:longitude;type:double;"`
	AirConditioned              *int64     `gorm:"column:air_conditioned;type:tinyint;"`
	DisabledAccess              *int64     `gorm:"column:disabled_access;type:tinyint;"`
	Tv                          *int64     `gorm:"column:tv;type:tinyint;"`
	Prestige                    *int64     `gorm:"column:prestige;type:tinyint;"`
	VeryHighPerfLabel           *int64     `gorm:"column:very_high_perf_label;type:tinyint;"`
	Loggia                      *int64     `gorm:"column:loggia;type:tinyint;"`
	Videocom                    *int64     `gorm:"column:videocom;type:tinyint;"`
	Cellar                      *int64     `gorm:"column:cellar;type:tinyint;"`
	Phl                         *int64     `gorm:"column:phl;type:tinyint;"`
	TaxExempBorlooLaw           *int64     `gorm:"column:tax_exemp_borloo_law;type:tinyint;"`
	Slice2                      *int64     `gorm:"column:slice2;type:tinyint;"`
	Garage                      *int64     `gorm:"column:garage;type:tinyint;"`
	Parking                     *int64     `gorm:"column:parking;type:tinyint;"`
	Box                         *int64     `gorm:"column:box;type:tinyint;"`
	SwimmingPool                *int64     `gorm:"column:swimming_pool;type:tinyint;"`
	Heating                     *int64     `gorm:"column:heating;type:tinyint;"`
	Terrace                     *int64     `gorm:"column:terrace;type:tinyint;"`
	GreenSpaces                 *int64     `gorm:"column:green_spaces;type:tinyint;"`
	EcoNeighborhood             *int64     `gorm:"column:eco_neighborhood;type:tinyint;"`
	Garden                      *int64     `gorm:"column:garden;type:tinyint;"`
	Keeper                      *int64     `gorm:"column:keeper;type:tinyint;"`
	FittedWardrobes             *int64     `gorm:"column:fitted_wardrobes;type:tinyint;"`
	ElectricShutters            *int64     `gorm:"column:electric_shutters;type:tinyint;"`
	TerraceArea                 *int64     `gorm:"column:terrace_area;type:int;"`
	DpeValueGaz                 *float64   `gorm:"column:dpe_value_gaz;type:double;"`
	TaxExempGirardinPaulLaw     *int64     `gorm:"column:tax_exemp_girardin_paul_law;type:tinyint;"`
	DpeValueConsumption         *float64   `gorm:"column:dpe_value_consumption;type:double;"`
	LeadCertificate             *int64     `gorm:"column:lead_certificate;type:tinyint;"`
	DiagnosisAsbestos           *int64     `gorm:"column:diagnosis_asbestos;type:tinyint;"`
	DpeSubmitted                *int64     `gorm:"column:dpe_submitted;type:tinyint;"`
	DpeVirgin                   *int64     `gorm:"column:dpe_virgin;type:tinyint;"`
	TaxExempPinel               *int64     `gorm:"column:tax_exemp_pinel;type:tinyint;"`
	GardenArea                  *int64     `gorm:"column:garden_area;type:int;"`
	Medias                      []*Media   `gorm:"foreignKey:ProgramID"`
	TaxExemption                string     `gorm:"column:tax_exemption;type:varchar;size:255;"`
	NbRooms                     string     `gorm:"column:nb_rooms;type:varchar;size:255;"`
	Developer                   Developer
	Properties                  []Property `gorm:"foreignKey:ProgramID"`
	Enabled                     int64      `gorm:"column:enabled;type:tinyint;"`
	PriceMax                    float64    `gorm:"column:price_max;type:double;"`
	NbRoomsMax                  int64      `gorm:"column:nb_rooms_max;type:int;"`
	AreaMax                     int64      `gorm:"column:area_max;type:int;"`
	PriceMin                    float64    `gorm:"column:price_min;type:double;"`
	NbRoomsMin                  int64      `gorm:"column:nb_rooms_min;type:int;"`
	AreaMin                     int64      `gorm:"column:area_min;type:int;"`
	Id                          int64      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
}

// TableName sets the insert table name for this struct type
func (p *Program) TableName() string {
	return "cdv_program"
}

// BeforeUpdate : process just before update
func (cp *Program) BeforeUpdate(tx *gorm.DB) (err error) {
	cp.UpdateAt = time.Now()
	return
}

// BeforeCreate : process just before create
func (cp *Program) BeforeCreate(tx *gorm.DB) (err error) {
	cp.UpdateAt = time.Now()
	cp.IntegrationDate = time.Now()
	cp.Enabled = int64(1)
	return
}
