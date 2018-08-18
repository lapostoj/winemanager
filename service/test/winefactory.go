package test

import (
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// AWine returns a wine instance to use in tests
func AWine() wine.Wine {
	storageLocation := wine.StorageLocation{
		Cellar:   "Cellar",
		Position: "3",
	}

	return wine.Wine{
		Name:            "Test Wine",
		Designation:     "Test Designation",
		Growth:          "Test Growth",
		Country:         "FR",
		Region:          "Bourgogne",
		Producer:        "Test Producer",
		Color:           wine.RED,
		Type:            wine.SEC,
		Year:            1963,
		Quantity:        3,
		Size:            wine.BOTTLE,
		StorageLocation: storageLocation,
		CreationTime:    time.Now().UTC(),
	}
}
