package request

import (
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
)

// PostBottleRequest defines the request that should be sent when posting a bottle
type PostBottleRequest struct {
	Year     int         `json:"year"`
	Size     bottle.Size `json:"size"`
	Quantity int         `json:"quantity"`
	CellarID int         `json:"cellarID"`
	Position string      `json:"position"`
}

// NewBottle transforms a PostBottleRequest into a Bottle
func (request PostBottleRequest) NewBottle() *bottle.Bottle {
	variation := bottle.Variation{
		Time:     time.Now().UTC(),
		Quantity: request.Quantity,
		Details:  "Adding bottle to stock",
	}

	return &bottle.Bottle{
		Year:     request.Year,
		Size:     request.Size,
		Quantity: request.Quantity,
		CellarID: request.CellarID,
		History:  []bottle.Variation{variation},
		StorageLocation: bottle.StorageLocation{
			Position: request.Position,
		},
	}
}
