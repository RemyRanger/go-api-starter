package repository

import (
	"gitlab.com/cdv-projects/go-apis/internal/services/flows"
	"gorm.io/gorm"
)

// SQLRepo : mysqlDB link
type SQLRepo struct {
	DB *gorm.DB
}

// NewSQLFlows : initializes a new FlowSqlRepo struct.
func NewSQLFlows(database *gorm.DB) flows.SQL {
	return &SQLRepo{
		DB: database,
	}
}
