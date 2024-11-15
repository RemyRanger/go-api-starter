package entities

type Query struct {
	OrderBy   string
	PageToken int64
	PageSize  int64
	SortDesc  bool
}

func NewQuery() *Query {
	return &Query{
		PageToken: 1,
		PageSize:  10,
		OrderBy:   "created_at",
		SortDesc:  false,
	}
}
