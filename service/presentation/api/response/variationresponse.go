package response

import "github.com/lapostoj/winemanager/service/domain/model/bottle"

// VariationResponse defines the object used when sending a variation.
type VariationResponse struct {
	Time     string `json:"time"`
	Quantity int    `json:"quantity"`
	Details  string `json:"details"`
}

// NewVariationResponse transforms a Variation in a VariationResponse.
func NewVariationResponse(variation bottle.Variation) *VariationResponse {
	return &VariationResponse{
		Time:     variation.Time.String(),
		Quantity: variation.Quantity,
		Details:  variation.Details,
	}
}

// NewVariationResponses transforms a slice of Variation in a slice of VariationResponse
func NewVariationResponses(variations []bottle.Variation) []VariationResponse {
	var variationResponses []VariationResponse
	for _, variation := range variations {
		variationResponses = append(variationResponses, *NewVariationResponse(variation))
	}
	return variationResponses
}
