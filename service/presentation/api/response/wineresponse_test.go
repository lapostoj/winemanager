package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/test"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewWineResponse(t *testing.T) {
	wine := test.AWineWithID()

	wineResponse := response.NewWineResponse(wine)

	assert.Equal(t, wineResponse.ID, wine.Key.ID)
	assert.Equal(t, wineResponse.Name, wine.Name)
	assert.Equal(t, wineResponse.Designation, wine.Designation)
	assert.Equal(t, wineResponse.Growth, wine.Growth)
	assert.Equal(t, wineResponse.Country, wine.Country)
	assert.Equal(t, wineResponse.Region, wine.Region)
	assert.Equal(t, wineResponse.Producer, wine.Producer)
	assert.Equal(t, wineResponse.Color, "RED")
	assert.Equal(t, wineResponse.Type, "SEC")
}

func TestNewWineResponses(t *testing.T) {
	wines := []wine.Wine{test.AWineWithID()}

	wineResponses := response.NewWineResponses(wines)

	assert.Equal(t, len(wineResponses), 1)
	assert.Equal(t, wineResponses[0].ID, wines[0].Key.ID)
	assert.Equal(t, wineResponses[0].Name, wines[0].Name)
	assert.Equal(t, wineResponses[0].Designation, wines[0].Designation)
	assert.Equal(t, wineResponses[0].Growth, wines[0].Growth)
	assert.Equal(t, wineResponses[0].Country, wines[0].Country)
	assert.Equal(t, wineResponses[0].Region, wines[0].Region)
	assert.Equal(t, wineResponses[0].Producer, wines[0].Producer)
	assert.Equal(t, wineResponses[0].Color, "RED")
	assert.Equal(t, wineResponses[0].Type, "SEC")
}
