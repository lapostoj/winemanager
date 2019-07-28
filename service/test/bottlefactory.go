package test

import (
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
)

// ABottle returns a bottle instance to use in tests
func ABottle() bottle.Bottle {
	storageLocation := bottle.StorageLocation{
		Position: "3",
	}

	return bottle.Bottle{
		Year:            1963,
		Size:            bottle.BOTTLE,
		Quantity:        3,
		CellarID:        123,
		WineID:          111,
		History:         []bottle.Variation{aVariation()},
		StorageLocation: storageLocation,
	}
}

func aVariation() bottle.Variation {
	return bottle.Variation{
		Time:     time.Now().UTC(),
		Quantity: 3,
		Details:  "Details",
	}
}
