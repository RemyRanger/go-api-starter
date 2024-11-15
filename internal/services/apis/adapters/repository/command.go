package repository

import (
	"context"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) UpsertApi(ctx context.Context, model *models.Api) (*models.Api, error) {
	if err := r.BeginWithCtx(ctx).Save(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Repository) UpsertApiWithTx(tx *gorm.DB, model *models.Api) (*models.Api, error) {
	if err := tx.Save(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Repository) DeleteApi(ctx context.Context, id uuid.UUID) error {
	model := &models.Api{
		ID: id,
	}
	if err := r.BeginWithCtx(ctx).Delete(&model).Error; err != nil {
		return err
	}

	return nil
}
