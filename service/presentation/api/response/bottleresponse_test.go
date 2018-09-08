package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/test"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewBottleResponse(t *testing.T) {
	bottle := test.ABottle()

	bottleResponse := response.NewBottleResponse(bottle)

	assert.Equal(t, bottleResponse.Year, bottle.Year)
	assert.Equal(t, bottleResponse.Size, "BOTTLE")
	assert.Equal(t, bottleResponse.Quantity, bottle.Quantity)
	assert.Equal(t, bottleResponse.CellarID, bottle.CellarID)
	assert.Equal(t, len(bottleResponse.History), len(bottle.History))
	assert.NotNil(t, bottle.Wine)
	assert.NotNil(t, bottleResponse.StorageLocation)
}
