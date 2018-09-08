package bottle

import "github.com/lapostoj/winemanager/service/domain/model/wine"

// Bottle defines the bottle object for our domain.
type Bottle struct {
	Year            int
	Size            Size
	Quantity        int
	CellarID        int
	Wine            wine.Wine
	History         []Variation
	StorageLocation StorageLocation
}
