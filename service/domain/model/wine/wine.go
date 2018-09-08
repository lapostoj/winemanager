package wine

// Wine defines the wine object for our domain.
type Wine struct {
	Name        string
	Designation string
	Growth      string
	Country     string
	Region      string
	Color       Color
	Type        Type
	Producer    string
}
