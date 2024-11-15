package repository

import (
	"context"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
)

func (r *Repository) FindApis(ctx context.Context, query *entities.Query) ([]models.Api, int64, error) {
	offset := (query.PageToken - 1) * query.PageSize
	orderQuery := ""
	if query.SortDesc {
		orderQuery = query.OrderBy + " desc"
	} else {
		orderQuery = query.OrderBy + " asc"
	}

	elems := []models.Api{}
	if err := r.BeginWithCtx(ctx).Offset(int(offset)).Limit(int(query.PageSize)).Order(orderQuery).Find(&elems).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := r.BeginWithCtx(ctx).Model(&models.Api{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return elems, count, nil
}

func (r *Repository) GetApi(ctx context.Context, id uuid.UUID) (*models.Api, error) {
	model := &models.Api{}
	if err := r.BeginWithCtx(ctx).First(model, id).Error; err != nil {
		return nil, err
	}

	return model, nil
}
