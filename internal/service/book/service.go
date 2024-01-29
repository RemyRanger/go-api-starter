package flows

import (
	"context"

	"gitlab.com/cdv-projects/go-apis/internal/common/entities"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
	"gitlab.com/cdv-projects/go-apis/internal/common/telemetry"
)

// ServiceName : service name
const ServiceName string = "flows"

// Service : interface for flow service
type Service interface {
	CreateFlow(context.Context, *models.Flow) (*models.Flow, error)
	UpdateFlow(context.Context, *models.Flow, string) (*models.Flow, error)
	UpdateFlowStatus(context.Context, bool, string, string) error

	GetFlow(context.Context, string) (*models.Flow, error)
	FindFlows(context.Context, *entities.Query) ([]*models.Flow, int64, error)
	FindFlowsOrigins(context.Context, *entities.QueryFlowsOrigins) ([]string, error)
	FindFlowsStatistics(context.Context) (*models.FlowsStatistics, error)
	FindJobsByFlowJobName(context.Context, string, *entities.Query) ([]*models.TalendFlowsStatistics, int64, error)
	FindContactsByFlowID(context.Context, string, *entities.Query) ([]*models.Contact, int64, error)
}

type flowService struct {
	sql SQL
}

// NewService : create service
func NewService(repo SQL) Service {
	return &flowService{
		sql: repo,
	}
}

// CreateFlow : add a flow
func (srv *flowService) CreateFlow(ctx context.Context, m *models.Flow) (*models.Flow, error) {
	flow, err := srv.sql.CreateFlow(ctx, m)
	if err != nil {
		telemetry.LogError(ServiceName, "Error create flow in database.", err)
		return nil, err
	}

	return flow, nil
}

// GetFlow : get a flow
func (srv *flowService) GetFlow(ctx context.Context, id string) (*models.Flow, error) {
	flow, err := srv.sql.GetFlow(ctx, id)
	if err != nil {
		telemetry.LogError(ServiceName, "Error get flow.", err)
		return nil, err
	}

	return flow, nil
}

// FindFlows : list all flows
func (srv *flowService) FindFlows(ctx context.Context, q *entities.Query) ([]*models.Flow, int64, error) {
	flows, count, err := srv.sql.FindFlows(ctx, q)
	if err != nil {
		telemetry.LogError(ServiceName, "Error find flows.", err)
		return nil, 0, err
	}

	return flows, count, nil
}

// FindFlowsOrigins : list all flows origins
func (srv *flowService) FindFlowsOrigins(ctx context.Context, q *entities.QueryFlowsOrigins) ([]string, error) {
	origins, err := srv.sql.FindFlowsOrigins(ctx, q)
	if err != nil {
		telemetry.LogError(ServiceName, "Error find flows origins.", err)
		return nil, err
	}

	return origins, nil
}

// UpdateFlow : update a flow
func (srv *flowService) UpdateFlow(ctx context.Context, m *models.Flow, id string) (*models.Flow, error) {
	flow, err := srv.sql.UpdateFlow(ctx, m, id)
	if err != nil {
		telemetry.LogError(ServiceName, "Error update flow.", err)
		return nil, err
	}

	return flow, nil
}

// FindFlowsStatistics : fetch global flow statistics
func (srv *flowService) FindFlowsStatistics(ctx context.Context) (*models.FlowsStatistics, error) {
	flowsStatistics, err := srv.sql.FindFlowsStatistics(ctx)
	if err != nil {
		telemetry.LogError(ServiceName, "Error find flows statistics.", err)
		return nil, err
	}

	return flowsStatistics, nil
}

// FindJobsByFlowJobName : list all flow jobs
func (srv *flowService) FindJobsByFlowJobName(ctx context.Context, jobName string, q *entities.Query) ([]*models.TalendFlowsStatistics, int64, error) {
	jobs, count, err := srv.sql.FindJobsByFlowJobName(ctx, jobName, q)
	if err != nil {
		telemetry.LogError(ServiceName, "Error find flow jobs.", err)
		return nil, 0, err
	}

	return jobs, count, nil
}

// UpdateFlowStatus update one flow visibility status
func (srv *flowService) UpdateFlowStatus(ctx context.Context, enable bool, subscriptionID, id string) error {
	if err := srv.sql.UpdateFlowStatus(ctx, enable, subscriptionID, id); err != nil {
		telemetry.LogError(ServiceName, "Error update flows status.", err)
		return err
	}

	return nil
}

// FindContactsByFlowID : get all contacts by flow id
func (srv *flowService) FindContactsByFlowID(ctx context.Context, id string, q *entities.Query) ([]*models.Contact, int64, error) {
	contacts, count, err := srv.sql.FindContactsByFlowID(ctx, id, q)
	if err != nil {
		telemetry.LogError(ServiceName, "Error find contacts.", err)
		return nil, 0, err
	}

	return contacts, count, nil
}
