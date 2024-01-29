package entities

// SearchListPrograms defines model for search_list_programs.
type SearchListPrograms struct {
	Filters *struct {
		// Filtres clef/valeur.
		KeyValue *[]struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"key_value,omitempty"`

		// Filtres Range.
		Range *[]struct {
			Key      string `json:"key"`
			ValueMax string `json:"value_max"`
			ValueMin string `json:"value_min"`
		} `json:"range,omitempty"`
	} `json:"filters,omitempty"`
	Geometries *struct {
		BoundingBox *BoundingBox `json:"bounding_box,omitempty"`
		Circle      *Circle      `json:"circle,omitempty"`
	} `json:"geometries,omitempty"`

	// Numéro de la page.
	Page *int64 `json:"page,omitempty"`

	// Nombre maximum d'éléments par page.
	PageLimit *int64 `json:"page_limit,omitempty"`
}

// BoundingBox defines model for bounding_box.
type BoundingBox struct {
	Bottom float64
	Left   float64
	Right  float64
	Top    float64
}

// Circle defines model for circle.
type Circle struct {
	Center *struct {
		Latitude  float64
		Longitude float64
	}
	Distance int64
}

// Polygon defines model for polygon.
type Polygon struct {
	Points []struct {
		Latitude  float64
		Longitude float64
	}
}
