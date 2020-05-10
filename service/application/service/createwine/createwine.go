package createwine

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// CreateWineService defines the interface for a CreateWineService
type CreateWineService interface {
	Execute(ctx context.Context, wine *wine.Wine) (int64, error)
}

// CreateWine implements a service to get Wines
type CreateWine struct {
	WineRepository wine.Repository
}

// Execute persists the passed wine and return its id.
func (service CreateWine) Execute(ctx context.Context, wine *wine.Wine) (int64, error) {
	return service.WineRepository.SaveWine(ctx, wine)
}
