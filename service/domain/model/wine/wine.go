package wine

import "cloud.google.com/go/datastore"

// Wine defines the wine object for our domain.
type Wine struct {
	Key         *datastore.Key `datastore:"__key__"`
	Name        string
	Designation string
	Growth      string
	Country     string
	Region      string
	Color       Color
	Type        Type
	Producer    string
}
