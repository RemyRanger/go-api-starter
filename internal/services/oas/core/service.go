package core

import (
	"bytes"
	"context"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"
	ports_api "go-api-starter/internal/services/apis/ports"
	ports_oas "go-api-starter/internal/services/oas/ports"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/uuid"
)

type Service struct {
	repositoryOas ports_oas.Repository
	repositoryApi ports_api.Repository
}

func New(repositoryOas ports_oas.Repository, repositoryApi ports_api.Repository) *Service {
	return &Service{
		repositoryOas: repositoryOas,
		repositoryApi: repositoryApi,
	}
}

func (s *Service) FindOas(ctx context.Context, apiId uuid.UUID, query *entities.Query) ([]models.Oas, int64, error) {
	return s.repositoryOas.FindOas(ctx, apiId, query)
}

func (s *Service) CreateOasFromFile(ctx context.Context, apiId uuid.UUID, buf *bytes.Buffer) (*models.Oas, error) {
	// Load & Validate OAS file
	loader := &openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromData(buf.Bytes())
	if err != nil {
		return nil, err
	}
	err = doc.Validate(ctx)
	if err != nil {
		return nil, err
	}

	// Convert file to []bytes
	data, err := doc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	// Save file as OAS object in db
	oas := &models.Oas{
		ApiID:          apiId,
		ApiVersion:     doc.Info.Version,
		OpenapiVersion: doc.OpenAPI,
		JsonSpec:       models.JSON(data),
	}

	tx, err := s.repositoryOas.BeginTransactionWithCtx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	modelOas, err := s.repositoryOas.UpsertOasWithTx(tx, apiId, oas)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Find linked API
	modelApi := &models.Api{}
	if modelApi, err = s.repositoryApi.GetApi(ctx, modelOas.ApiID); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Active new Openapi specification on Api
	modelApi.ActiveOasID = &modelOas.ID
	if _, err := s.repositoryApi.UpsertApiWithTx(tx, modelApi); err != nil {
		tx.Rollback()
		return nil, err
	}

	return modelOas, tx.Commit().Error
}

func (s *Service) GetOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) (*models.Oas, error) {
	return s.repositoryOas.GetOas(ctx, apiId, oasId)
}

func (s *Service) DeleteOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) error {
	return s.repositoryOas.DeleteOas(ctx, apiId, oasId)
}
