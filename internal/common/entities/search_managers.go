package entities

import (
	"net/http"
)

// SearchManagers defines model for SearchManagers.
type SearchManagers struct {
	PageSize     *int
	PageToken    *int
	QueryName    *string
	FilterEnable *bool
}

// Bind : just a post-process after a decode..
func (sm *SearchManagers) Bind(r *http.Request) error {
	return nil
}
