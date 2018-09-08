package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/test"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewWineResponse(t *testing.T) {
	wine := test.AWine()

	wineResponse := response.NewWineResponse(wine)

	assert.Equal(t, wineResponse.Name, wine.Name)
	assert.Equal(t, wineResponse.Designation, wine.Designation)
	assert.Equal(t, wineResponse.Growth, wine.Growth)
	assert.Equal(t, wineResponse.Country, wine.Country)
	assert.Equal(t, wineResponse.Region, wine.Region)
	assert.Equal(t, wineResponse.Producer, wine.Producer)
	assert.Equal(t, wineResponse.Color, "RED")
	assert.Equal(t, wineResponse.Type, "SEC")
}
