package cellar

import (
	"context"
)

// Repository is the interface for a repository in the persistence layer for the Cellar aggregate
type Repository interface {
	SaveCellar(ctx context.Context, cellar *Cellar) (string, error)
}
