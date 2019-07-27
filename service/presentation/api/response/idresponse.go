package response

// IDResponse defines the object used when sending a cellar.
type IDResponse struct {
	ID string `json:"ID"`
}

// NewIDResponse transforms a string in an IDResponse.
func NewIDResponse(id string) *IDResponse {
	return &IDResponse{
		ID: id,
	}
}
