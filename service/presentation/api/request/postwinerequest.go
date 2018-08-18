package request

import (
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// PostWineRequest defines the request that should be sent when posting a wine
type PostWineRequest struct {
	Name        string     `json:"name"`
	Designation string     `json:"designation"`
	Growth      string     `json:"growth"`
	Year        int        `json:"year"`
	Country     string     `json:"country"`
	Region      string     `json:"region"`
	Color       wine.Color `json:"color"`
	Type        wine.Type  `json:"type"`
	Quantity    int        `json:"quantity"`
	Producer    string     `json:"producer"`
	Size        wine.Size  `json:"size"`
	Cellar      string     `json:"cellar"`
	Position    string     `json:"position"`
}

// NewWine transforms a PostWineRequest into a Wine
func (request PostWineRequest) NewWine() *wine.Wine {
	return &wine.Wine{
		Name:        request.Name,
		Designation: request.Designation,
		Growth:      request.Growth,
		Year:        request.Year,
		Country:     request.Country,
		Region:      request.Region,
		Color:       request.Color,
		Type:        request.Type,
		Quantity:    request.Quantity,
		Producer:    request.Producer,
		Size:        request.Size,
		History:     *new([]wine.Variation),
		StorageLocation: wine.StorageLocation{
			Cellar:   request.Cellar,
			Position: request.Position,
		},
		CreationTime: time.Now().UTC(),
	}
}
