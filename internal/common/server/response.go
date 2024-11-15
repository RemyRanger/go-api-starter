package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

const fieldErrMsg = "Field validation failed on the '%s' tag"

type ErrorResponse struct {
	ErrorMessage   string         `json:"error"`
	Details        []ErrorDetails `json:"details,omitempty"`
	HTTPStatusCode int            `json:"status"`
}

type ErrorDetails struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func RespondError(w http.ResponseWriter, err error, code int) {
	log.Error().Stack().Err(errors.WithStack(err)).Msg("Error during process")

	if code == http.StatusBadRequest {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorDetails, len(ve))
			for i, fe := range ve {
				out[i] = ErrorDetails{fe.Field(), fmt.Sprintf(fieldErrMsg, fe.Tag())}
			}

			RespondWithBody(w, &ErrorResponse{
				HTTPStatusCode: code,
				ErrorMessage:   http.StatusText(code),
				Details:        out, // only show details for StatusBadRequest
			}, code)
		}
	} else {
		RespondWithBody(w, &ErrorResponse{
			HTTPStatusCode: code,
			ErrorMessage:   http.StatusText(code),
		}, code)
	}
}

func RespondWithBody(w http.ResponseWriter, respData any, code int) {
	jsonResponse, jsonError := json.Marshal(&respData)
	if jsonError != nil {
		RespondError(w, jsonError, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
