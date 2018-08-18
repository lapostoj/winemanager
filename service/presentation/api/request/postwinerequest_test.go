package request_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/wine"

	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/stretchr/testify/assert"
)

func TestNewWine(t *testing.T) {
	request := request.PostWineRequest{
		Name:        "Wine",
		Designation: "Designation",
		Growth:      "Growth",
		Year:        2018,
		Country:     "FR",
		Region:      "Region",
		Color:       wine.RED,
		Type:        wine.DOUX,
		Quantity:    1,
		Producer:    "Producer",
		Size:        wine.BOTTLE,
		Cellar:      "Cellar",
		Position:    "3",
	}

	wine := request.NewWine()

	assert.Equal(t, request.Name, wine.Name)
	assert.Equal(t, request.Designation, wine.Designation)
	assert.Equal(t, request.Growth, wine.Growth)
	assert.Equal(t, request.Year, wine.Year)
	assert.Equal(t, request.Country, wine.Country)
	assert.Equal(t, request.Region, wine.Region)
	assert.Equal(t, request.Color, wine.Color)
	assert.Equal(t, request.Type, wine.Type)
	assert.Equal(t, request.Quantity, wine.Quantity)
	assert.Equal(t, request.Producer, wine.Producer)
	assert.Equal(t, request.Size, wine.Size)
	assert.Equal(t, request.Cellar, wine.StorageLocation.Cellar)
	assert.Equal(t, request.Position, wine.StorageLocation.Position)
	assert.Equal(t, 0, len(wine.History))
	assert.NotNil(t, wine.CreationTime)
}
