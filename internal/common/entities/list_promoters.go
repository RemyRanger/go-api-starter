package entities

// ListPromoters defines model for listPromoters.
type ListPromoters struct {
	Pagination *struct {
		Page     *int64 `json:"page,omitempty"`
		PageSize *int64 `json:"page_size,omitempty"`
		Total    *int64 `json:"total,omitempty"`
	} `json:"pagination,omitempty"`
	Promoters *[]struct {
		Name *string `json:"name,omitempty"`
	} `json:"promoters,omitempty"`
}
