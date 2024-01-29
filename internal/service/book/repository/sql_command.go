package repository

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
)

// CreateFlow : create flow in database
func (r *SQLRepo) CreateFlow(ctx context.Context, record *models.Flow) (*models.Flow, error) {
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateFlow : create flow in database
func (r *SQLRepo) UpdateFlow(ctx context.Context, record *models.Flow, id string) (*models.Flow, error) {
	if err := r.DB.Model(&models.Flow{}).Where("uuid = ?", id).Updates(record).Take(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateSCPIStatus : update scpi status
func (r *SQLRepo) UpdateFlowStatus(ctx context.Context, enable bool, subscriptionID, id string) error {
	// Search if scpi visibility already exist for request apikey
	qoVisibility := r.DB.Model(&models.FlowVisibility{}).
		Where("flow_uuid = ?", id).
		Where("subscription_uuid = ?", subscriptionID)
	flowVisibility := &models.FlowVisibility{}
	if err := qoVisibility.Find(flowVisibility).Error; err != nil {
		return err
	}

	if flowVisibility.SubscriptionUUID != uuid.Nil { // If flow already has visibility, update it
		updateFlowVisibility := make(map[string]any) // GORM will ignore the value like 0, nil, “” or false https://medium.com/@speedforcerun/golang-how-to-solve-gorm-not-updating-data-when-you-may-have-none-zero-field-724ccb351e8b
		if enable {
			updateFlowVisibility["enabled"] = 1
		} else {
			updateFlowVisibility["enabled"] = 0
		}

		if err := qoVisibility.Updates(&updateFlowVisibility).Error; err != nil {
			return err
		}
	} else { // Else, create it
		flowUUID, _ := uuid.Parse(id)
		subscriptionUUID, _ := uuid.Parse(subscriptionID)
		flowVisibility := &models.FlowVisibility{
			FlowUUID:         flowUUID,
			SubscriptionUUID: subscriptionUUID,
		}

		if enable {
			flowVisibility.Enabled = 1
		} else {
			flowVisibility.Enabled = 0
		}

		if err := r.DB.Create(flowVisibility).Error; err != nil {
			return err
		}
	}

	return nil
}
