package ports

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
)

// FLOW

// UUID : UUID for copier mapping
func (f *Flow) UUID(uuid uuid.UUID) {
	UUID := uuid.String()
	f.ID = &UUID
}

// CONTACT

// UUID : UUID for copier mapping
func (c *Contact) UUID(uuid uuid.UUID) {
	UUID := uuid.String()
	c.ID = &UUID
}

// ProjectUUID : ProjectUUID for copier mapping
func (c *Contact) ProjectUUID(uuid *uuid.UUID) {
	if uuid != nil {
		projectID := uuid.String()
		c.ProjectID = &projectID
	}
}

// FlowUUID : FlowUUID for copier mapping
func (c *Contact) FlowUUID(uuid *uuid.UUID) {
	if uuid != nil {
		flowID := uuid.String()
		c.FlowID = &flowID
	}
}

// ModelSQLToJob : convert TalendFlowsStatistics to Job
func ModelSQLToJob(talendFlowsStatistic *models.TalendFlowsStatistics) *Job {
	job := &Job{}

	copier.Copy(&job, &talendFlowsStatistic)

	if talendFlowsStatistic.TalendFlowsLog != nil {
		jobError := &JobError{}

		jobError.Message = talendFlowsStatistic.TalendFlowsLog.Message
		jobError.Origin = talendFlowsStatistic.TalendFlowsLog.Origin
		jobError.Type = talendFlowsStatistic.TalendFlowsLog.Type

		job.Error = jobError
	}

	return job
}
