package wine

import "time"

// Wine defines the wine object for our domain.
type Wine struct {
	Name            string
	Designation     string
	Growth          string
	Year            int
	Country         string
	Region          string
	Color           Color
	Type            Type
	Quantity        int
	Producer        string
	Size            Size
	History         []Variation
	StorageLocation StorageLocation
	CreationTime    time.Time
}

// NewWine creates a Wine struct with default values.
func NewWine() *Wine {
	wine := Wine{
		CreationTime: time.Now().UTC(),
	}
	return &wine
}
