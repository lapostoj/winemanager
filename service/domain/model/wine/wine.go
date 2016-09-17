package wine

import "time"

// Wine defines the wine object for our domain.
type Wine struct {
	Name            string          `json:"name"`
	Designation     string          `json:"designation"`
	Growth          string          `json:"growth"`
	Year            int             `json:"year"`
	Country         string          `json:"country"`
	Region          string          `json:"region"`
	Color           Color           `json:"color"`
	Type            Type            `json:"type"`
	Quantity        int             `json:"quantity"`
	Producer        string          `json:"producer"`
	Size            Size            `json:"size"`
	History         []Variation     `json:"history"`
	StorageLocation StorageLocation `json:"storage_location"`
	CreationTime    time.Time       `json:"creation_time"`
}

// NewWine creates a Wine struct with default values.
func NewWine() *Wine {
	wine := Wine{CreationTime: time.Now().UTC()}
	return &wine
}
