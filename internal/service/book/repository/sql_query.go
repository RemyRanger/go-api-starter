package repository

import (
	"context"

	"gitlab.com/cdv-projects/go-apis/internal/common/entities"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
	"gorm.io/plugin/dbresolver"
)

// GetFlow : find flow in database
func (r *SQLRepo) GetFlow(ctx context.Context, id string) (*models.Flow, error) {
	flow := &models.Flow{}
	if err := r.DB.Model(models.Flow{}).First(flow, "uuid = ?", id).Error; err != nil {
		return nil, err
	}
	return flow, nil
}

// FindFlows : list all flows in database
func (r *SQLRepo) FindFlows(ctx context.Context, q *entities.Query) ([]*models.Flow, int64, error) {
	size := 10
	offset := 0
	if q.PageSize != nil {
		size = *q.PageSize
	}
	if q.PageToken != nil {
		offset = (*q.PageToken - 1) * size
	}

	orderQuery := ""
	if q.OrderBy != nil && q.SortDirDesc != nil {
		if *q.SortDirDesc {
			orderQuery = *q.OrderBy + " desc"
		} else {
			orderQuery = *q.OrderBy + " asc"
		}
	}

	flows := []*models.Flow{}
	if err := r.DB.Offset(offset).Limit(size).Order(orderQuery).Find(&flows).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := r.DB.Model(&models.Flow{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return flows, count, nil
}

// FindFlowsOrigins : list all flows in database
func (r *SQLRepo) FindFlowsOrigins(ctx context.Context, q *entities.QueryFlowsOrigins) ([]string, error) {
	query := r.DB.Model(&models.Flow{})

	if q.TypeOffers != nil {
		query.Where("type_offers = ?", q.TypeOffers)
	}

	if q.Status != nil {
		query.Where("status = ?", q.Status)
	}

	origins := []string{}
	if err := query.Select("flow_origin").Find(&origins).Error; err != nil {
		return nil, err
	}

	return origins, nil
}

// FindFlowsStatistics : find all stats about flow jobs
func (r *SQLRepo) FindFlowsStatistics(ctx context.Context) (*models.FlowsStatistics, error) {
	var countAds int64
	if err := r.DB.Clauses(dbresolver.Use("secondary")).Model(&models.Property{}).Count(&countAds).Error; err != nil {
		return nil, err
	}

	var countFlows int64
	if err := r.DB.Model(&models.Flow{}).Count(&countFlows).Error; err != nil {
		return nil, err
	}

	var countActiveFlows int64
	if err := r.DB.Model(&models.Flow{}).Where("status = ?", "ENABLED").Count(&countActiveFlows).Error; err != nil {
		return nil, err
	}

	talendFlowsStatistics := []*models.TalendFlowsStatistics{}
	if err := r.DB.Clauses(dbresolver.Use("secondary")).
		Raw(`SELECT * FROM 
			(SELECT MAX(moment) as maxmoment, job FROM talend_flow_stats GROUP BY job) maxtime 
			INNER JOIN talend_flow_stats AS stats ON stats.job=maxtime.job 
			AND stats.moment=maxtime.maxmoment 
			WHERE origin IS NULL 
			AND message_type = 'end'`).Scan(&talendFlowsStatistics).Error; err != nil {
		return nil, err
	}

	var jobsOk []string
	var jobsError []string
	for _, talendFlowsStatistic := range talendFlowsStatistics {
		if talendFlowsStatistic.Message == "success" {
			jobsOk = append(jobsOk, talendFlowsStatistic.Job)
		}
		if talendFlowsStatistic.Message == "failure" {
			jobsError = append(jobsError, talendFlowsStatistic.Job)
		}
	}

	return &models.FlowsStatistics{
		NbAds:        countAds,
		NbFlows:      countActiveFlows,
		NbFlowsTotal: &countFlows,
		JobsOk:       jobsOk,
		JobsError:    jobsError,
	}, nil
}

// FindJobsByFlowJobName : list all flows in database
func (r *SQLRepo) FindJobsByFlowJobName(ctx context.Context, jobName string, q *entities.Query) ([]*models.TalendFlowsStatistics, int64, error) {
	size := 10
	offset := 0
	if q.PageSize != nil {
		size = *q.PageSize
	}
	if q.PageToken != nil {
		offset = (*q.PageToken - 1) * size
	}

	talendFlowsStatistics := []*models.TalendFlowsStatistics{}
	if err := r.DB.Clauses(dbresolver.Use("secondary")).Order("moment desc").
		Offset(offset).Limit(size).Where("origin IS NULL").
		Where(&models.TalendFlowsStatistics{MessageType: "end", Job: jobName}).Find(&talendFlowsStatistics).Error; err != nil {
		return nil, 0, err
	}

	for _, talendFlowsStatistic := range talendFlowsStatistics {
		if talendFlowsStatistic.Message == "failure" {
			talendFlowsLog := models.TalendFlowsLog{}
			if err := r.DB.Clauses(dbresolver.Use("secondary")).Where("pid = ?", talendFlowsStatistic.Pid).Find(&talendFlowsLog).Error; err != nil {
				return nil, 0, err
			}

			talendFlowsStatistic.TalendFlowsLog = &talendFlowsLog
		}
	}

	var count int64
	if err := r.DB.Clauses(dbresolver.Use("secondary")).Model(&models.TalendFlowsStatistics{}).
		Where("origin IS NULL").Where(&models.TalendFlowsStatistics{MessageType: "end", Job: jobName}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return talendFlowsStatistics, count, nil
}

// FindContactsByFlowID : find all contacts by flow_id in database
func (r *SQLRepo) FindContactsByFlowID(ctx context.Context, id string, q *entities.Query) ([]*models.Contact, int64, error) {
	size := 10
	offset := 0
	if q.PageSize != nil {
		size = *q.PageSize
	}
	if q.PageToken != nil {
		offset = (*q.PageToken - 1) * size
	}

	orderQuery := ""
	if q.OrderBy != nil && q.SortDirDesc != nil {
		if *q.SortDirDesc {
			orderQuery = *q.OrderBy + " desc"
		} else {
			orderQuery = *q.OrderBy + " asc"
		}
	}

	contacts := []*models.Contact{}
	if err := r.DB.Offset(offset).Limit(size).Order(orderQuery).Where("flow_uuid = ?", id).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := r.DB.Model(&models.Contact{}).Where("flow_uuid = ?", id).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return contacts, count, nil
}
