package flows

import (
	"context"

	"gitlab.com/cdv-projects/go-apis/internal/common/entities"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
)

// SQL : interface for flow repository
type SQL interface {
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
