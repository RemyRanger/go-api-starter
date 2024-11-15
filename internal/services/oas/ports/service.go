package ports

import (
	"bytes"
	"context"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
)

// Service : service interface
type Service interface {
	FindOas(ctx context.Context, apiId uuid.UUID, query *entities.Query) ([]models.Oas, int64, error)
	CreateOasFromFile(ctx context.Context, apiId uuid.UUID, buf *bytes.Buffer) (*models.Oas, error)
	GetOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) (*models.Oas, error)
	DeleteOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) error
}
