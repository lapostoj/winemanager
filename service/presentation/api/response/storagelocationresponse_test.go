package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewStorageLocationResponse(t *testing.T) {
	storageLocation := bottle.StorageLocation{
		Position: "3",
	}

	storageLocationResponse := response.NewStorageLocationResponse(storageLocation)

	assert.Equal(t, storageLocationResponse.Position, storageLocation.Position)
}
