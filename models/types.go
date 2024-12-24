// models/types.go
package models

// CityData represents the complete data structure for a city
type CityData struct {
	CdMun      string     `json:"cd_mun"`
	Name       string     `json:"name"`
	Bounds     Bounds     `json:"bounds"`
	Categories []Category `json:"categories"`
}

// Bounds represents the geographical boundaries of a city
type Bounds struct {
	Type       string     `json:"type"`
	Properties BoundProps `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

// BoundProps contains the properties of the boundary feature
type BoundProps struct {
	Name    string  `json:"name"`
	AreaKm2 float64 `json:"area_km2"`
}

// Geometry represents the GeoJSON geometry
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// Category represents a group of related layers
type Category struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Icon   string  `json:"icon"`
	Layers []Layer `json:"layers"`
}

// Layer represents a single data layer
type Layer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
