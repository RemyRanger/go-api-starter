package ports

import (
	"context"

	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindApis(ctx context.Context, query *entities.Query) ([]models.Api, int64, error)

	UpsertApi(ctx context.Context, model *models.Api) (*models.Api, error)
	UpsertApiWithTx(tx *gorm.DB, model *models.Api) (*models.Api, error)

	GetApi(ctx context.Context, apiId uuid.UUID) (*models.Api, error)

	DeleteApi(ctx context.Context, apiId uuid.UUID) error
}
