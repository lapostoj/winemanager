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
		Country:     "FR",
		Region:      "Region",
		Color:       wine.RED,
		Type:        wine.DOUX,
		Producer:    "Producer",
	}

	wine := request.NewWine()

	assert.Equal(t, request.Name, wine.Name)
	assert.Equal(t, request.Designation, wine.Designation)
	assert.Equal(t, request.Growth, wine.Growth)
	assert.Equal(t, request.Country, wine.Country)
	assert.Equal(t, request.Region, wine.Region)
	assert.Equal(t, request.Color, wine.Color)
	assert.Equal(t, request.Type, wine.Type)
	assert.Equal(t, request.Producer, wine.Producer)
}
