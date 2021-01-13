// Package railroad provides useful types.
package railroad

// FeatureCollection corresponds a top level object.
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

// Feature corresponds each entries of feature collection.
type Feature struct {
	Type     string   `json:"type"`
	Property Property `json:"properties"`
	Geometry Geometry `json:"geometry"`
}

// Property holds properties about station.
type Property struct {
	RailroadType string `json:"鉄道区分"`
	BusinessType string `json:"事業者種別"`
	Line         string `json:"路線名"`
	Company      string `json:"運営会社"`
	Station      string `json:"駅名"`
}

// Geometry holds coordinates slice.
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates []Coordinates `json:"coordinates"`
}

// Coordinates is a pair of longitude and latitude.
type Coordinates []float64
