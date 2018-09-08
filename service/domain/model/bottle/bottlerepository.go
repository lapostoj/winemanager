package bottle

import (
	"context"
)

// Repository is the interface for a repository in the persistence layer for the Bottle aggregate
type Repository interface {
	SaveBottle(ctx context.Context, bottle *Bottle) (string, error)
}
