package response

// IDResponse defines the object used when sending a cellar.
type IDResponse struct {
	ID int64 `json:"ID"`
}

// NewIDResponse transforms a string in an IDResponse.
func NewIDResponse(id int64) *IDResponse {
	return &IDResponse{
		ID: id,
	}
}
