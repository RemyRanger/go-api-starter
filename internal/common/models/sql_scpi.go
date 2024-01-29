package models

import (
	"time"

	"gorm.io/gorm"
)

// Scpi : database scpi entity model
type Scpi struct {
	UpdateAt                  time.Time `gorm:"column:update_at;type:datetime;"`
	PerformanceReferenceDate  time.Time `gorm:"column:performance_reference_date;type:datetime;"`
	WithdrawalValueUpdateDate time.Time `gorm:"column:withdrawal_value_update_date;type:datetime;"`
	CreateAt                  time.Time `gorm:"column:create_at;type:datetime;"`
	ScpiCreationDate          time.Time `gorm:"column:scpi_creation_date;type:datetime;"`
	ScpiStatutoryTerm         time.Time `gorm:"column:scpi_statutory_term;type:datetime;"`
	NominalValueUpdateDate    time.Time `gorm:"column:nominal_value_update_date;type:datetime;"`
	DistributionReferenceDate time.Time `gorm:"column:distribution_reference_date;type:datetime;"`
	WithdrawalDate            time.Time `gorm:"column:withdrawal_date;type:datetime;"`
	Name                      string    `gorm:"column:name;type:varchar;size:255;"`
	TypeScpi                  string    `gorm:"column:type;type:varchar;size:255;"`
	FlowOrigin                string    `gorm:"column:flow_origin;type:varchar;size:255;"`
	ScpiMinimumDuration       string    `gorm:"column:scpi_minimum_duration;type:text;size:65535;"`
	IncomePaymentDates        string    `gorm:"column:income_payment_dates;type:varchar;size:1024;"`
	ScpiMinimumRisk           string    `gorm:"column:scpi_minimum_risk;type:text;size:65535;"`
	ScpiMinimumKnowledge      string    `gorm:"column:scpi_minimum_knowledge;type:text;size:65535;"`
	GeoDistribution           string    `gorm:"column:geo_distribution;type:text;size:65535;"`
	UrlLogo                   string    `gorm:"column:url_logo;type:varchar;size:1024;"`
	FinancialDistribution     string    `gorm:"column:financial_distribution;type:text;size:65535;"`
	Subtype                   string    `gorm:"column:subtype;type:varchar;size:255;"`
	Strategy                  string    `gorm:"column:strategy;type:text;size:65535;"`
	VacancyRate               float64   `gorm:"column:vacancy_rate;type:double;"`
	NbSites                   float64   `gorm:"column:nb_sites;type:double;"`
	Tof                       float64   `gorm:"column:tof;type:double;"`
	SubscriptionPrice         float64   `gorm:"column:subscription_price;type:double;"`
	SubscriptionFees          float64   `gorm:"column:subscription_fees;type:double;"`
	NominalValue              float64   `gorm:"column:nominal_value;type:double;"`
	Vpm                       float64   `gorm:"column:vpm;type:double;"`
	Bonus                     float64   `gorm:"column:bonus;type:double;"`
	Income                    float64   `gorm:"column:income;type:double;"`
	WithdrawalValue           float64   `gorm:"column:withdrawal_value;type:double;"`
	Tdvm                      float64   `gorm:"column:tdvm;type:double;"`
	Tri15                     float64   `gorm:"column:tri15;type:double;"`
	ManagementFees            float64   `gorm:"column:management_fees;type:double;"`
	Tri10                     float64   `gorm:"column:tri10;type:double;"`
	Tri5                      float64   `gorm:"column:tri5;type:double;"`
	SubscriptionMinValue      float64   `gorm:"column:subscription_min_value;type:double;"`
	Area                      float64   `gorm:"column:area;type:double;"`
	SubscriptionMinPiece      int64     `gorm:"column:subscription_min_piece;type:int;"`
	EnjoymentDelay            int64     `gorm:"column:enjoyment_delay;type:int;"`
	BestRetentionPeriod       int64     `gorm:"column:best_retention_period;type:int;"`
	Risk                      int64     `gorm:"column:risk;type:int;"`
	Capital                   int64     `gorm:"column:capital;type:int;"`
	ManagerId                 int64     `gorm:"column:manager_id;type:int;"`
	NbLeases                  float64   `gorm:"column:nb_leases;type:double;"`
	Top                       float64   `gorm:"column:top;type:double;"`
	Id                        int64     `gorm:"primary_key;column:id;type:int;"`
}

func (s *Scpi) BeforeCreate(tx *gorm.DB) (err error) {
	s.UpdateAt = time.Now()
	s.CreateAt = time.Now()
	return
}

func (s *Scpi) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdateAt = time.Now()
	return
}

// TableName sets the insert table name for this struct type
func (c *Scpi) TableName() string {
	return "cdv_scpi"
}
