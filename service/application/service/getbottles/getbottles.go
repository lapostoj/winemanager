package getbottles

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
)

// GetBottlesService defines the interface for a GetBottlesService
type GetBottlesService interface {
	ForCellarID(ctx context.Context, accountID int) ([]bottle.Bottle, error)
}

// GetBottles implements a service to get Bottles
type GetBottles struct {
	BottleRepository bottle.Repository
}

// ForCellarID returns the bottles with the provided cellarID.
func (service GetBottles) ForCellarID(ctx context.Context, cellarID int) ([]bottle.Bottle, error) {
	var bottles []bottle.Bottle
	err := service.BottleRepository.FindBottlesForCellarID(ctx, &bottles, cellarID)

	return bottles, err
}
