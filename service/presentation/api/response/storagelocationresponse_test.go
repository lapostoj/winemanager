package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/wine"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewStorageLocationResponse(t *testing.T) {
	storageLocation := wine.StorageLocation{
		Cellar:   "Cellar",
		Position: "3",
	}

	storageLocationResponse := response.NewStorageLocationResponse(storageLocation)

	assert.Equal(t, storageLocationResponse.Cellar, storageLocation.Cellar)
	assert.Equal(t, storageLocationResponse.Position, storageLocation.Position)
}
