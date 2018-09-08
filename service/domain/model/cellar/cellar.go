package cellar

import "time"

// Cellar defines the cellar object for our domain.
type Cellar struct {
	Name         string
	AccountID    int
	CreationTime time.Time
}

// NewCellar creates a Cellar struct with default values.
func NewCellar() *Cellar {
	return &Cellar{
		CreationTime: time.Now().UTC(),
	}
}
