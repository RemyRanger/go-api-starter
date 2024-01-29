package models

import "github.com/olivere/elastic"

// Cluster: Cluster struct
type ClusterES struct {
	Count       *int64   `json:"count,omitempty"`
	Lat         *float64 `json:"lat,omitempty"`
	Lng         *float64 `json:"lng,omitempty"`
	Maxlat      *float64 `json:"maxlat,omitempty"`
	Maxlng      *float64 `json:"maxlng,omitempty"`
	Minlat      *float64 `json:"minlat,omitempty"`
	Minlng      *float64 `json:"minlng,omitempty"`
	Price       *int64   `json:"price,omitempty"`
	ItemIds     *[]int64 `json:"itemIds,omitempty"`
	SinglePoint *bool    `json:"singlePoint,omitempty"`
}

// MapAggregationValues : map aggregation value to ProgramCluster field
func (c *ClusterES) MapAggregationValues(bucket *elastic.AggregationBucketKeyItem) {

	lng, _ := bucket.Avg("lng")
	c.Lng = lng.Value

	lat, _ := bucket.Avg("lat")
	c.Lat = lat.Value

	c.Count = &bucket.DocCount

	minlat, _ := bucket.Min("minlat")
	c.Minlat = minlat.Value

	minlng, _ := bucket.Min("minlng")
	c.Minlng = minlng.Value

	maxlat, _ := bucket.Max("maxlat")
	c.Maxlat = maxlat.Value

	maxlng, _ := bucket.Max("maxlng")
	c.Maxlng = maxlng.Value

	price, _ := bucket.Min("price")
	if price.Value != nil {
		priceInt64 := int64(*price.Value)
		c.Price = &priceInt64
	}

	itemIds, _ := bucket.Terms("itemIds")
	var itemIdsSlice []int64
	for _, bucket := range itemIds.Buckets {
		value := bucket.Key.(float64)
		itemIdsSlice = append(itemIdsSlice, int64(value))
	}
	c.ItemIds = &itemIdsSlice

	clusterLatitudes, _ := bucket.Terms("clusterLatitudes")
	singlePoint := len(clusterLatitudes.Buckets) <= 1
	c.SinglePoint = &singlePoint

}
