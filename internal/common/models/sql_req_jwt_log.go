package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ReqJWTLog : database ReqJWTLog entity model
type ReqJWTLog struct {
	CreatedAt           time.Time
	HeaderAgent         string `gorm:"size:255"`
	Proto               string `gorm:"size:255"`
	CompanyCode         string `gorm:"size:255"`
	Host                string `gorm:"size:255"`
	HeaderOrigin        string `gorm:"size:255"`
	Method              string `gorm:"size:255"`
	HeaderAuthorization string `gorm:"size:255"`
	ApikeyId            string `gorm:"size:255;type:varchar"`
	CompanyId           string `gorm:"size:255;type:varchar"`
	SubscriptionId      string `gorm:"size:255;type:varchar"`
	ContentLength       int64
	UUID                uuid.UUID `gorm:"primary_key;type:varchar"`
}

// TableName sets the insert table name for this struct type
func (r *ReqJWTLog) TableName() string {
	return "req_jwt_logs"
}

// BeforeCreate : This functions are called before creating Base
func (r *ReqJWTLog) BeforeCreate(tx *gorm.DB) (err error) {
	r.UUID = uuid.New()
	return
}
