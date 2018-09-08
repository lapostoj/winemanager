package bottle

import "time"

// Variation defines a variation in the stock.
// Details could contain the origin (gift, place where bought)
// or the person it was drank with.
type Variation struct {
	Time     time.Time
	Quantity int
	Details  string
}
