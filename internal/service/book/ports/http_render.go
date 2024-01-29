package ports

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gitlab.com/cdv-projects/go-apis/internal/services/flows"
)

var validate *validator.Validate

// Bind : just a post-process after a decode..
func (nf *NewFlow) Bind(r *http.Request) error {
	// Validation
	validate = validator.New()
	if err := validate.Struct(nf); err != nil {
		log.Debug().Str("service", flows.ServiceName).Err(err).Msg("Error validating incoming flow.")
		return err
	}

	return nil
}

// Bind : just a post-process after a decode..
func (si *StatusIn) Bind(r *http.Request) error {
	// Validation
	validate = validator.New()
	if err := validate.Struct(si); err != nil {
		log.Debug().Str("service", flows.ServiceName).Err(err).Msg("Error validating incoming visibility status.")
		return err
	}

	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (f *Flow) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (fs *FlowsStatistics) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (rfu *ResponseFetchFlows) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (f *ResponseFetchFlowsOrigins) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (fj *FlowJobs) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (p *ResponseFetchContacts) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
