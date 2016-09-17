package wine

import "time"

// Variation defines a variation in the stock.
// Details could contain the origin (gift, place where bought)
// or the person it was drank with.
type Variation struct {
	Date     time.Time `json:"date"`
	Quantity int       `json:"quantity"`
	Details  string    `json:"details"`
}
