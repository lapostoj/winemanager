package response

import "github.com/lapostoj/winemanager/service/domain/model/bottle"

// StorageLocationResponse defines the object used when sending a storage location.
type StorageLocationResponse struct {
	Position string `json:"position"`
}

// NewStorageLocationResponse transforms a StorageLocation in a StorageLocationResponse.
func NewStorageLocationResponse(storageLocation bottle.StorageLocation) *StorageLocationResponse {
	return &StorageLocationResponse{
		Position: storageLocation.Position,
	}
}
