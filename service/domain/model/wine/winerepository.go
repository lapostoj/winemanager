package wine

import (
	"context"
)

// Repository is the interface for a repository in the persistence layer for the Wine normalized value
type Repository interface {
	GetWines(ctx context.Context, wines *[]Wine) error
	SaveWine(ctx context.Context, wine *Wine) (string, error)
}
