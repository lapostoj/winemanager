package response

import "github.com/lapostoj/winemanager/service/domain/model/wine"

// StorageLocationResponse defines the object used when sending a storage location.
type StorageLocationResponse struct {
	Cellar   string `json:"cellar"`
	Position string `json:"position"`
}

// NewStorageLocationResponse transforms a StorageLocation in a StorageLocationResponse.
func NewStorageLocationResponse(storageLocation wine.StorageLocation) *StorageLocationResponse {
	return &StorageLocationResponse{
		Cellar:   storageLocation.Cellar,
		Position: storageLocation.Position,
	}
}
