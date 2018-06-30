package response

import "github.com/lapostoj/winemanager/service/domain/model/wine"

// WineResponse defines the object used when sending a wine.
type WineResponse struct {
	Name            string                  `json:"name"`
	Designation     string                  `json:"designation"`
	Growth          string                  `json:"growth"`
	Year            int                     `json:"year"`
	Country         string                  `json:"country"`
	Region          string                  `json:"region"`
	Color           string                  `json:"color"`
	Type            string                  `json:"type"`
	Quantity        int                     `json:"quantity"`
	Producer        string                  `json:"producer"`
	Size            string                  `json:"size"`
	History         []VariationResponse     `json:"history"`
	StorageLocation StorageLocationResponse `json:"storageLocation"`
}

// NewWineResponse transforms a Wine in WineReponse
func NewWineResponse(wine wine.Wine) *WineResponse {
	return &WineResponse{
		Name:            wine.Name,
		Designation:     wine.Designation,
		Growth:          wine.Growth,
		Year:            wine.Year,
		Country:         wine.Country,
		Region:          wine.Region,
		Color:           wine.Color.String(),
		Type:            wine.Type.String(),
		Quantity:        wine.Quantity,
		Producer:        wine.Producer,
		Size:            wine.Size.String(),
		History:         NewVariationResponses(wine.History),
		StorageLocation: *NewStorageLocationResponse(wine.StorageLocation),
	}
}

// NewWineResponses transforms a slice of Wine in a slice of WineReponse
func NewWineResponses(wines []wine.Wine) []WineResponse {
	wineResponses := []WineResponse{}
	for _, wine := range wines {
		wineResponses = append(wineResponses, *NewWineResponse(wine))
	}
	return wineResponses
}
