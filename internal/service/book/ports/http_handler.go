package ports

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	"gitlab.com/cdv-projects/go-apis/internal/common/entities"
	cdvErrors "gitlab.com/cdv-projects/go-apis/internal/common/errors"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
	"gitlab.com/cdv-projects/go-apis/internal/services/flows"
)

// HTTPHandler : handler http type
type HTTPHandler struct {
	svc flows.Service
}

// NewHTTPHandler : create new http handler
func NewHTTPHandler(svc flows.Service) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

// FindFlows : Fetch Flows
// (GET /flows)
func (h *HTTPHandler) FindFlows(w http.ResponseWriter, r *http.Request, params FindFlowsParams) {
	q := &entities.Query{}
	/* #nosec */
	copier.Copy(&q, &params)

	flows, count, err := h.svc.FindFlows(r.Context(), q)
	if err != nil {
		render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		return
	}

	respData := []Flow{}
	/* #nosec */
	copier.Copy(&respData, &flows)

	respDataFormated := ResponseFetchFlows{
		Count: &count,
		Flows: &respData,
	}

	render.Status(r, http.StatusOK)
	/* #nosec */
	render.Render(w, r, &respDataFormated)
}

// CreateFlow : Create Flow
// (POST /flows)
func (h *HTTPHandler) CreateFlow(w http.ResponseWriter, r *http.Request) {
	reqData := &NewFlow{}
	if err := render.Bind(r, reqData); err != nil {
		render.Render(w, r, cdvErrors.ErrInvalidRequest(err))
		return
	}

	flow := &models.Flow{}
	/* #nosec */
	copier.Copy(&flow, &reqData)

	flow, err := h.svc.CreateFlow(r.Context(), flow)
	if err != nil {
		render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		return
	}

	respData := &Flow{}
	/* #nosec */
	copier.Copy(&respData, &flow)

	render.Status(r, http.StatusCreated)
	/* #nosec */
	render.Render(w, r, respData)
}

// Fetch Flows origins
// (GET /flows/origins)
func (h *HTTPHandler) FindFlowsOrigins(w http.ResponseWriter, r *http.Request, params FindFlowsOriginsParams) {
	q := &entities.QueryFlowsOrigins{}
	/* #nosec */
	copier.Copy(&q, &params)
	origins, err := h.svc.FindFlowsOrigins(r.Context(), q)
	if err != nil {
		render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		return
	}

	respDataFormated := ResponseFetchFlowsOrigins{
		FlowsOrigins: &origins,
	}

	render.Status(r, http.StatusOK)
	/* #nosec */
	render.Render(w, r, &respDataFormated)
}

// Fetch global Flows statistics
// (GET /flows/stats)
func (h *HTTPHandler) FindFlowsStatistics(w http.ResponseWriter, r *http.Request) {
	flowsStats, err := h.svc.FindFlowsStatistics(r.Context())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Render(w, r, cdvErrors.ErrDataNotFound(err, cdvErrors.ErrNotFound))
		} else {
			render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		}
		return
	}
	respData := &FlowsStatistics{}
	/* #nosec */
	copier.Copy(&respData, &flowsStats)

	// set cache
	w.Header().Set("Etag", uuid.NewString())
	w.Header().Set("Cache-Control", "max-age=3600") // 1h

	render.Status(r, http.StatusOK)
	/* #nosec */
	render.Render(w, r, respData)
}

// Returns jobs by flow job name code
// (GET /flows/{job_name}/jobs)
func (h *HTTPHandler) FindJobsByFlowJobName(w http.ResponseWriter, r *http.Request, jobName string, params FindJobsByFlowJobNameParams) {
	q := &entities.Query{}
	/* #nosec */
	copier.Copy(&q, &params)

	jobs, count, err := h.svc.FindJobsByFlowJobName(r.Context(), jobName, q)
	if err != nil {
		render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		return
	}

	respData := []Job{}
	for _, job := range jobs {
		respJob := ModelSQLToJob(job)
		respData = append(respData, *respJob)
	}

	respDataFormated := FlowJobs{
		Count:    &count,
		FlowJobs: &respData,
	}

	render.Status(r, http.StatusOK)
	/* #nosec */
	render.Render(w, r, &respDataFormated)
}

// FindFlowByID : Returns flow by ID
// (GET /flows/{id})
func (h *HTTPHandler) FindFlowByID(w http.ResponseWriter, r *http.Request, id string) {
	flow, err := h.svc.GetFlow(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Render(w, r, cdvErrors.ErrDataNotFound(err, cdvErrors.ErrNotFound))
		} else {
			render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		}
		return
	}
	respData := &Flow{}
	/* #nosec */
	copier.Copy(&respData, &flow)

	render.Status(r, http.StatusOK)
	/* #nosec */
	render.Render(w, r, respData)
}

// UpdateFlow : Update Flow
// (PUT /flows/{id})
func (h *HTTPHandler) UpdateFlow(w http.ResponseWriter, r *http.Request, id string) {
	reqData := &NewFlow{}
	if err := render.Bind(r, reqData); err != nil {
		render.Render(w, r, cdvErrors.ErrInvalidRequest(err))
		return
	}

	flow := &models.Flow{}
	/* #nosec */
	copier.Copy(&flow, &reqData)

	flow, err := h.svc.UpdateFlow(r.Context(), flow, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Render(w, r, cdvErrors.ErrDataNotFound(err, cdvErrors.ErrNotFound))
		} else {
			render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		}
		return
	}

	respData := &Flow{}
	/* #nosec */
	copier.Copy(&respData, &flow)

	render.Status(r, http.StatusCreated)
	/* #nosec */
	render.Render(w, r, respData)
}

// Update Flow Visibility
// (PUT /flows/{id}/visibility)
func (h *HTTPHandler) UpdateFlowVisibility(w http.ResponseWriter, r *http.Request, id string) {
	reqData := &StatusIn{}
	if err := render.Bind(r, reqData); err != nil {
		render.Render(w, r, cdvErrors.ErrInvalidRequest(err))
		return
	}

	err := h.svc.UpdateFlowStatus(r.Context(), *reqData.Enable, reqData.SubscriptionID, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Render(w, r, cdvErrors.ErrDataNotFound(err, cdvErrors.ErrNotFound))
		} else {
			render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		}
		return
	}

	render.NoContent(w, r)
}

// Fetch Contacts by flow ID
// (GET /flows/{id}/contacts)
func (h *HTTPHandler) FindContactsByFlowID(w http.ResponseWriter, r *http.Request, id string, params FindContactsByFlowIDParams) {
	q := &entities.Query{}
	/* #nosec */
	copier.Copy(&q, &params)

	contacts, count, err := h.svc.FindContactsByFlowID(r.Context(), id, q)
	if err != nil {
		render.Render(w, r, cdvErrors.ErrInternal(err, cdvErrors.ErrInternalProcess))
		return
	}

	respData := []Contact{}
	/* #nosec */
	copier.Copy(&respData, &contacts)

	respDataFormated := ResponseFetchContacts{
		Count:    &count,
		Contacts: &respData,
	}

	render.Status(r, http.StatusOK)
	/* #nosec */
	render.Render(w, r, &respDataFormated)
}
