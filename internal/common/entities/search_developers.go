package entities

import (
	"net/http"
)

// SearchDevelopers defines model for SearchDevelopers.
type SearchDevelopers struct {
	PageSize     *int
	PageToken    *int
	QueryName    *string
	FilterEnable *bool
}

// Bind : just a post-process after a decode..
func (sm *SearchDevelopers) Bind(r *http.Request) error {
	return nil
}
