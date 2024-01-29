package models

import (
	"time"
)

// TalendFlowsStatistics : database flow entity model
type TalendFlowsStatistics struct {
	Moment          time.Time
	TalendFlowsLog  *TalendFlowsLog `gorm:"-"`
	JobRepositoryId string          `gorm:"size:255"`
	RootPid         string          `gorm:"size:20"`
	Message         string          `gorm:"size:255"`
	Project         string          `gorm:"size:50"`
	JobVersion      string          `gorm:"size:255"`
	Pid             string          `gorm:"size:20"`
	FatherPid       string          `gorm:"size:20"`
	Context         string          `gorm:"size:50"`
	Origin          string          `gorm:"size:255"`
	MessageType     string          `gorm:"size:255"`
	Job             string          `gorm:"size:255"`
	SystemPid       int64
	Duration        int64
}

// FlowsStatistics defines model for FlowsStatistics.
type TalendFlowsLog struct {
	Moment  time.Time
	Type    string `gorm:"size:255"`
	Message string `gorm:"size:255"`
	Job     string `gorm:"size:255"`
	Pid     string `gorm:"size:20"`
	Origin  string `gorm:"size:255"`
}

// FlowsStatistics defines model for FlowsStatistics.
type FlowsStatistics struct {
	NbFlowsTotal *int64
	JobsError    []string
	JobsOk       []string
	NbAds        int64
	NbFlows      int64
}

// TableName sets the insert table name for this struct type
func (f *TalendFlowsStatistics) TableName() string {
	return "talend_flow_stats"
}

// TableName sets the insert table name for this struct type
func (f *TalendFlowsLog) TableName() string {
	return "talend_flow_logs"
}
