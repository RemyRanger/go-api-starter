package ports

import (
	"context"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindOas(ctx context.Context, apiId uuid.UUID, query *entities.Query) ([]models.Oas, int64, error)

	GetOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) (*models.Oas, error)

	UpsertOas(ctx context.Context, apiId uuid.UUID, model *models.Oas) (*models.Oas, error)
	UpsertOasWithTx(tx *gorm.DB, apiId uuid.UUID, model *models.Oas) (*models.Oas, error)

	DeleteOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) error

	BeginTransactionWithCtx(ctx context.Context) (*gorm.DB, error)
}
