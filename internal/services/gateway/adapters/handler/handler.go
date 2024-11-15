package handler

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleGateway(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
