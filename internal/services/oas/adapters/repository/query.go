package repository

import (
	"context"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
)

func (r *Repository) FindOas(ctx context.Context, apiId uuid.UUID, query *entities.Query) ([]models.Oas, int64, error) {
	offset := (query.PageToken - 1) * query.PageSize
	orderQuery := ""
	if query.SortDesc {
		orderQuery = query.OrderBy + " desc"
	} else {
		orderQuery = query.OrderBy + " asc"
	}

	elems := []models.Oas{}
	if err := r.BeginWithCtx(ctx).Offset(int(offset)).Limit(int(query.PageSize)).Order(orderQuery).Where("api_id = ?", apiId).Find(&elems).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := r.BeginWithCtx(ctx).Model(&models.Oas{}).Where("api_id = ?", apiId).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return elems, count, nil
}

func (r *Repository) GetOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) (*models.Oas, error) {
	model := &models.Oas{}
	if err := r.BeginWithCtx(ctx).Where("api_id = ?", apiId).First(model, oasId).Error; err != nil {
		return nil, err
	}

	return model, nil
}
