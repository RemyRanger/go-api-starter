package core

import (
	"context"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"
	"go-api-starter/internal/services/apis/ports"

	"github.com/google/uuid"
)

type Service struct {
	repository ports.Repository
}

func New(repository ports.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) FindApis(ctx context.Context, query *entities.Query) ([]models.Api, int64, error) {
	return s.repository.FindApis(ctx, query)
}

func (s *Service) CreateApi(ctx context.Context, model *models.Api) (*models.Api, error) {
	return s.repository.UpsertApi(ctx, model)
}

func (s *Service) GetApi(ctx context.Context, id uuid.UUID) (*models.Api, error) {
	return s.repository.GetApi(ctx, id)
}

func (s *Service) UpdateApi(ctx context.Context, id uuid.UUID, model *models.Api) (*models.Api, error) {
	model.ID = id
	return s.repository.UpsertApi(ctx, model)
}

func (s *Service) DeleteApi(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteApi(ctx, id)
}
