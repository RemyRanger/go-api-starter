package entities

import (
	"net/http"
)

// SearchManagersCera defines model for SearchManagersCera.
type SearchManagersCera struct {
	Name       *string
	Email      *string
	Zip        *string
	Siret      *string
	ScpiStatus *string
}

// Bind : just a post-process after a decode..
func (sm *SearchManagersCera) Bind(r *http.Request) error {
	return nil
}
