package request

import (
	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// PostWineRequest defines the request that should be sent when posting a wine
type PostWineRequest struct {
	Name        string     `json:"name"`
	Designation string     `json:"designation"`
	Growth      string     `json:"growth"`
	Country     string     `json:"country"`
	Region      string     `json:"region"`
	Producer    string     `json:"producer"`
	Color       wine.Color `json:"color"`
	Type        wine.Type  `json:"type"`
}

// NewWine transforms a PostWineRequest into a Wine
func (request PostWineRequest) NewWine() *wine.Wine {
	return &wine.Wine{
		Name:        request.Name,
		Designation: request.Designation,
		Growth:      request.Growth,
		Country:     request.Country,
		Region:      request.Region,
		Producer:    request.Producer,
		Color:       request.Color,
		Type:        request.Type,
	}
}
