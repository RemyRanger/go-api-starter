package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"go-api-starter/internal/common/entities"
	"go-api-starter/internal/common/server"
	"go-api-starter/internal/services/oas/ports"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
)

const handler_name = "oas-handler"

var tracer = otel.Tracer("oas-handler-tracer")

type Handler struct {
	service ports.Service
}

func NewHandler(service ports.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Find Oas
// (GET /apis/{api_id}/oas)
func (h *Handler) FindOas(w http.ResponseWriter, r *http.Request, apiId ports.Id, params ports.FindOasParams) {
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
	modelsFound, total, err := h.service.FindOas(ctx, apiId, query)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Map to response body
	var elems []ports.Oas
	if err := copier.Copy(&elems, &modelsFound); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	respBody := ports.OasList{
		Page: ports.Page{
			Total:    total,
			Page:     query.PageToken,
			PageSize: query.PageSize,
		},
		Elems: elems,
	}

	server.RespondWithBody(w, &respBody, http.StatusOK)
}

// Upload Oas file.
// (POST /apis/{api_id}/oas/_upload)
func (h *Handler) UploadOas(w http.ResponseWriter, r *http.Request, apiId ports.Id) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Extract Multipart file from request
	file, _, err := r.FormFile("oas")
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Copy file to bytes buffer
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Debug().Err(err).Msg("")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Call service
	modelCreated, err := h.service.CreateOasFromFile(ctx, apiId, buf)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	// Map model to response body
	respBody := new(ports.Oas)
	if err := copier.Copy(&respBody, &modelCreated); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, &respBody, http.StatusOK)
}

// Delete Oas.
// (DELETE /apis/{api_id}/oas/{oas_id})
func (h *Handler) DeleteOas(w http.ResponseWriter, r *http.Request, apiId ports.Id, oasId ports.Id) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Call service
	if err := h.service.DeleteOas(ctx, apiId, oasId); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Oas.
// (GET /apis/{api_id}/oas/{oas_id})
func (h *Handler) GetOas(w http.ResponseWriter, r *http.Request, apiId ports.Id, oasId ports.Id) {
	// Start otlp tracer
	ctx, span := tracer.Start(r.Context(), handler_name)
	defer span.End()

	// Call service
	modelFound, err := h.service.GetOas(ctx, apiId, oasId)
	if err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	if modelFound == nil {
		server.RespondError(w, fmt.Errorf("entity not found with id %s", apiId), http.StatusFound)
		return
	}

	// Map model to response
	respData := new(ports.Oas)
	if err := copier.Copy(&respData, &modelFound); err != nil {
		server.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	server.RespondWithBody(w, &respData, http.StatusOK)
}

// Download Oas file.
// (GET /apis/{api_id}/oas/{oas_id}/_download)
func (h *Handler) DownloadOas(w http.ResponseWriter, r *http.Request, apiId ports.Id, oasId ports.Id) {
	w.WriteHeader(http.StatusNotImplemented)
}
