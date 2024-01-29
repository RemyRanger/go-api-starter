package entities

// QueryMetrics entity.
type QueryMetrics struct {
	NbDays         *int64
	SubscriptionId string
}

// NewQueryMetrics : create new QueryMetrics entity
func NewQueryMetrics() *QueryMetrics {
	assignNbDays := int64(1)

	return &QueryMetrics{
		NbDays: &assignNbDays,
	}
}
