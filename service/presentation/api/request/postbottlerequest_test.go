package request_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/stretchr/testify/assert"
)

func TestNewBottle(t *testing.T) {
	request := request.PostBottleRequest{
		Year:     2018,
		Quantity: 1,
		Size:     bottle.BOTTLE,
		CellarID: 123,
		Position: "3",
	}

	bottle := request.NewBottle()

	assert.Equal(t, request.Year, bottle.Year)
	assert.Equal(t, request.Quantity, bottle.Quantity)
	assert.Equal(t, request.Size, bottle.Size)
	assert.Equal(t, request.CellarID, bottle.CellarID)
	assert.Equal(t, request.Position, bottle.StorageLocation.Position)
	assert.Equal(t, 1, len(bottle.History))
	assert.Equal(t, request.Quantity, bottle.History[0].Quantity)
	assert.NotNil(t, bottle.History[0].Time)
	assert.NotNil(t, bottle.History[0].Details)
}
