package entities

import (
	"net/http"
)

// SearchSCPIs defines model for SearchSCPIs.
type SearchSCPIs struct {
	PageSize     *int
	PageToken    *int
	QueryName    *string
	ManagerId    *string
	ScpiTypes    *[]string
	FilterEnable *bool
}

// Bind : just a post-process after a decode..
func (ss *SearchSCPIs) Bind(r *http.Request) error {
	return nil
}
