package response

import (
	"github.com/lapostoj/winemanager/service/domain/model/bottle"
)

// BottleResponse defines the object used when sending a bottle.
type BottleResponse struct {
	Year            int                     `json:"year"`
	Size            string                  `json:"size"`
	Quantity        int                     `json:"quantity"`
	CellarID        int                     `json:"cellarID"`
	Wine            WineResponse            `json:"wine"`
	History         []VariationResponse     `json:"history"`
	StorageLocation StorageLocationResponse `json:"storageLocation"`
}

// NewBottleResponse transforms a Bottle in BottleReponse
func NewBottleResponse(bottle bottle.Bottle) *BottleResponse {
	return &BottleResponse{
		Year:            bottle.Year,
		Size:            bottle.Size.String(),
		Quantity:        bottle.Quantity,
		CellarID:        bottle.CellarID,
		Wine:            *NewWineResponse(bottle.Wine),
		History:         NewVariationResponses(bottle.History),
		StorageLocation: *NewStorageLocationResponse(bottle.StorageLocation),
	}
}

// NewBottleResponses transforms a slice of Bottle in a slice of BottleReponse
func NewBottleResponses(bottles []bottle.Bottle) []BottleResponse {
	bottleResponses := []BottleResponse{}
	for _, bottle := range bottles {
		bottleResponses = append(bottleResponses, *NewBottleResponse(bottle))
	}
	return bottleResponses
}
