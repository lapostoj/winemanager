package response

import "github.com/lapostoj/winemanager/service/domain/model/wine"

// WineResponse defines the object used when sending a wine.
type WineResponse struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Growth      string `json:"growth"`
	Country     string `json:"country"`
	Region      string `json:"region"`
	Color       string `json:"color"`
	Type        string `json:"type"`
	Producer    string `json:"producer"`
}

// NewWineResponse transforms a Wine in WineReponse
func NewWineResponse(wine wine.Wine) *WineResponse {
	return &WineResponse{
		ID:          wine.Key.ID,
		Name:        wine.Name,
		Designation: wine.Designation,
		Growth:      wine.Growth,
		Country:     wine.Country,
		Region:      wine.Region,
		Color:       wine.Color.String(),
		Type:        wine.Type.String(),
		Producer:    wine.Producer,
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
