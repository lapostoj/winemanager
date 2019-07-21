package createcellar

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/cellar"
)

// CreateCellarService defines the interface for a CreateCellarService
type CreateCellarService interface {
	Execute(ctx context.Context, cellar *cellar.Cellar) (string, error)
}

// CreateCellar implements a service to get Cellars
type CreateCellar struct {
	CellarRepository cellar.Repository
}

// Execute returns the cellars with the provided accountID.
func (service CreateCellar) Execute(ctx context.Context, cellar *cellar.Cellar) (string, error) {
	return service.CellarRepository.SaveCellar(ctx, cellar)
}
