package response

import "github.com/lapostoj/winemanager/service/domain/model/cellar"

// CellarResponse defines the object used when sending a cellar.
type CellarResponse struct {
	ID           int64  `json:"ID"`
	Name         string `json:"name"`
	AccountID    int    `json:"accountID"`
	CreationTime string `json:"creationTime"`
}

// NewCellarResponse transforms a Cellar in a CellarResponse.
func NewCellarResponse(cellar cellar.Cellar) *CellarResponse {
	return &CellarResponse{
		ID:           cellar.Key.ID,
		Name:         cellar.Name,
		AccountID:    cellar.AccountID,
		CreationTime: cellar.CreationTime.String(),
	}
}

// NewCellarResponses transforms a slice of Cellar in a slice of CellarResponse
func NewCellarResponses(cellars []cellar.Cellar) []CellarResponse {
	var cellarResponses []CellarResponse
	for _, cellar := range cellars {
		cellarResponses = append(cellarResponses, *NewCellarResponse(cellar))
	}
	return cellarResponses
}
