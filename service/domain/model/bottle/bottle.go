package bottle

// Bottle defines the bottle object for our domain.
type Bottle struct {
	Year            int
	Size            Size
	Quantity        int
	CellarID        int64
	WineID          int64
	History         []Variation
	StorageLocation StorageLocation
}
