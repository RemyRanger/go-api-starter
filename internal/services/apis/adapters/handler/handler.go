package handler

import (
	"encoding/json"
	"fmt"
	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/models"
	"go-api-starter/internal/common/server"
	"go-api-starter/internal/services/apis/ports"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
)

const handler_name = "apis-handler"

var tracer = otel.Tracer("apis-handler-tracer")

type Handler struct {
	service ports.Service
}

func NewHandler(service ports.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Find APIs
// (GET /apis)
func (h *Handler) FindApis(w http.ResponseWriter, r *http.Request, params ports.FindApisParams) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Validate query params
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		server.RespondError(w, err, http.StatusBadRequest)
		return
	}

	// Map request query to entity
	query := entities.NewQuery()
	if err := copier.CopyWithOption(&query, &params, copier.Option{IgnoreEmpty: true}); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Call service
	modelsFound, total, err := h.service.FindApis(ctx, query)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Map to response body
	var elems []ports.Api
	if err := copier.Copy(&elems, &modelsFound); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	respBody := ports.ApiList{
		Page: ports.Page{
			Total:    total,
			Page:     query.PageToken,
			PageSize: query.PageSize,
		},
		Elems: elems,
	}

	server.RespondWithBody(w, &respBody, http.StatusOK)
}

// Create API
// (POST /apis)
func (h *Handler) CreateApi(w http.ResponseWriter, r *http.Request) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Decode & Validate request body
	var body ports.ApiIn
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		server.RespondError(w, err, http.StatusBadRequest)
		return
	}

	// Map request body to model
	var model models.Api
	if err := copier.Copy(&model, &body); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Call service
	modelCreated, err := h.service.CreateApi(ctx, &model)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Map model to response body
	respBody := new(ports.Api)
	if err := copier.Copy(&respBody, &modelCreated); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, &respBody, http.StatusOK)
}

// Delete API.
// (DELETE /apis/{api_id})
func (h *Handler) DeleteApi(w http.ResponseWriter, r *http.Request, apiId ports.Id) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Call service
	if err := h.service.DeleteApi(ctx, apiId); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get API.
// (GET /apis/{api_id})
func (h *Handler) GetApi(w http.ResponseWriter, r *http.Request, apiId ports.Id) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Call service
	modelFound, err := h.service.GetApi(ctx, apiId)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	if modelFound == nil {
		server.RespondError(w, fmt.Errorf("api not found with id %s", apiId), http.StatusFound)
		return
	}

	// Map model to response
	respData := new(ports.Api)
	if err := copier.Copy(&respData, &modelFound); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, &respData, http.StatusOK)
}

// Update API.
// (PUT /apis/{api_id})
func (h *Handler) UpdateApi(w http.ResponseWriter, r *http.Request, apiId ports.Id) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Decode & Validate request body
	var body ports.ApiIn
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		server.RespondError(w, err, http.StatusBadRequest)
		return
	}

	// Map request body to model
	var model models.Api
	if err := copier.Copy(&model, &body); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Call service
	modelUpdated, err := h.service.UpdateApi(ctx, apiId, &model)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Map model to response body
	var respBody ports.Api
	if err := copier.Copy(&respBody, &modelUpdated); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, &respBody, http.StatusOK)
}
