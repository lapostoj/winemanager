package response

import "github.com/lapostoj/winemanager/service/domain/model/wine"

// VariationResponse defines the object used when sending a variation.
type VariationResponse struct {
	Date     string `json:"date"`
	Quantity int    `json:"quantity"`
	Details  string `json:"details"`
}

// NewVariationResponse transforms a Varition in a VariationResponse.
func NewVariationResponse(variation wine.Variation) *VariationResponse {
	return &VariationResponse{
		Date:     variation.Date.String(),
		Quantity: variation.Quantity,
		Details:  variation.Details,
	}
}

// NewVariationResponses transforms a slice of Variation in a slice of VariationResponse
func NewVariationResponses(variations []wine.Variation) []VariationResponse {
	var variationResponses []VariationResponse
	for _, variation := range variations {
		variationResponses = append(variationResponses, *NewVariationResponse(variation))
	}
	return variationResponses
}
