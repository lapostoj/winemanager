package createbottle

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
)

// CreateBottleService defines the interface for a CreateBottleService
type CreateBottleService interface {
	Execute(ctx context.Context, bottle *bottle.Bottle) (string, error)
}

// CreateBottle implements a service to get Bottles
type CreateBottle struct {
	BottleRepository bottle.Repository
}

// Execute persists the passed bottle and return its id.
func (service CreateBottle) Execute(ctx context.Context, bottle *bottle.Bottle) (string, error) {
	return service.BottleRepository.SaveBottle(ctx, bottle)
}
