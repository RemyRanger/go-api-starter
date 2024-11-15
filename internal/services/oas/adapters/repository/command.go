package repository

import (
	"context"
	"go-api-starter/internal/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) UpsertOas(ctx context.Context, apiId uuid.UUID, model *models.Oas) (*models.Oas, error) {
	if err := r.BeginWithCtx(ctx).Save(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Repository) UpsertOasWithTx(tx *gorm.DB, apiId uuid.UUID, model *models.Oas) (*models.Oas, error) {
	if err := tx.Save(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Repository) DeleteOas(ctx context.Context, apiId uuid.UUID, oasId uuid.UUID) error {
	model := &models.Oas{
		ID:    oasId,
		ApiID: apiId,
	}
	if err := r.BeginWithCtx(ctx).Delete(&model).Error; err != nil {
		return err
	}

	return nil
}
