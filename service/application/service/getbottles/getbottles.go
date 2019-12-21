package getbottles

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

// GetBottlesService defines the interface for a GetBottlesService
type GetBottlesService interface {
	ForCellarID(ctx context.Context, accountID int) ([]response.BottleResponse, error)
}

// GetBottles implements a service to get Bottles
type GetBottles struct {
	BottleRepository bottle.Repository
	WineRepository   wine.Repository
}

// ForCellarID returns the bottles with the provided cellarID.
func (service GetBottles) ForCellarID(ctx context.Context, cellarID int) ([]response.BottleResponse, error) {
	var bottles []bottle.Bottle
	err := service.BottleRepository.FindBottlesForCellarID(ctx, &bottles, cellarID)
	var wines []wine.Wine
	for _, bottle := range bottles {
		var wine wine.Wine
		service.WineRepository.GetWineByID(ctx, bottle.WineID, &wine)
		wines = append(wines, wine)
	}

	return response.NewBottleResponses(bottles, wines), err
}
