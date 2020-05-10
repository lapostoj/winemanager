package wine

import (
	"context"
)

// Repository is the interface for a repository in the persistence layer for the Wine normalized value
type Repository interface {
	SaveWine(ctx context.Context, wine *Wine) (int64, error)
	GetWines(ctx context.Context, wines *[]Wine) error
	GetWineByID(ctx context.Context, ID int64, wine *Wine) error
}
