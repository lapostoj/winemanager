package wine

import (
	"context"
)

// Repository is the interface for a repository in the persistence layer for the Wine aggregate
type Repository interface {
	SaveTestWine(ctx context.Context) error
	GetWinesInStock(ctx context.Context, wines *[]Wine) error
	SaveWine(ctx context.Context, wine *Wine) (string, error)
}