package ports

import (
	"context"

	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
)

// Service : service interface
type Service interface {
	FindApis(context.Context, *entities.Query) ([]models.Api, int64, error)
	CreateApi(context.Context, *models.Api) (*models.Api, error)
	GetApi(context.Context, uuid.UUID) (*models.Api, error)
	UpdateApi(context.Context, uuid.UUID, *models.Api) (*models.Api, error)
	DeleteApi(context.Context, uuid.UUID) error
}
