package response_test

import (
	"testing"
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewVariationResponse(t *testing.T) {
	variation := bottle.Variation{
		Time:     time.Now().UTC(),
		Quantity: 3,
		Details:  "with friends",
	}

	variationResponse := response.NewVariationResponse(variation)

	assert.Equal(t, variationResponse.Time, variation.Time.String())
	assert.Equal(t, variationResponse.Quantity, variation.Quantity)
	assert.Equal(t, variationResponse.Details, variation.Details)
}

func TestNewVariationResponses(t *testing.T) {
	variations := []bottle.Variation{
		bottle.Variation{
			Time:     time.Now().UTC(),
			Quantity: 3,
			Details:  "with friends",
		},
	}

	variationResponses := response.NewVariationResponses(variations)

	assert.Equal(t, len(variationResponses), 1)
	assert.Equal(t, variationResponses[0].Time, variations[0].Time.String())
	assert.Equal(t, variationResponses[0].Quantity, variations[0].Quantity)
	assert.Equal(t, variationResponses[0].Details, variations[0].Details)
}
