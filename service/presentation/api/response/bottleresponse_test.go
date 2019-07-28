package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/test"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewBottleResponse(t *testing.T) {
	bottle := test.ABottle()
	wine := test.AWineWithID()

	bottleResponse := response.NewBottleResponse(bottle, wine)

	assert.Equal(t, bottleResponse.Year, bottle.Year)
	assert.Equal(t, bottleResponse.Size, "BOTTLE")
	assert.Equal(t, bottleResponse.Quantity, bottle.Quantity)
	assert.Equal(t, bottleResponse.CellarID, bottle.CellarID)
	assert.Equal(t, bottleResponse.Wine.ID, wine.Key.ID)
	assert.Equal(t, len(bottleResponse.History), len(bottle.History))
	assert.NotNil(t, bottleResponse.StorageLocation)
}

func TestNewBottleResponses(t *testing.T) {
	bottles := []bottle.Bottle{test.ABottle()}
	wines := []wine.Wine{test.AWineWithID()}

	bottleResponses := response.NewBottleResponses(bottles, wines)

	assert.Equal(t, len(bottleResponses), 1)
	assert.Equal(t, bottleResponses[0].Year, bottles[0].Year)
	assert.Equal(t, bottleResponses[0].Size, "BOTTLE")
	assert.Equal(t, bottleResponses[0].Quantity, bottles[0].Quantity)
	assert.Equal(t, bottleResponses[0].CellarID, bottles[0].CellarID)
	assert.Equal(t, bottleResponses[0].Wine.ID, wines[0].Key.ID)
	assert.Equal(t, len(bottleResponses[0].History), len(bottles[0].History))
	assert.NotNil(t, bottleResponses[0].StorageLocation)
}
