package getwines

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// GetWinesService defines the interface for a GetWinesService
type GetWinesService interface {
	Execute(ctx context.Context) ([]wine.Wine, error)
}

// GetWines implements a service to get Wines
type GetWines struct {
	WineRepository wine.Repository
}

// Execute returns all the stored wines
func (service GetWines) Execute(ctx context.Context) ([]wine.Wine, error) {
	var wines []wine.Wine
	err := service.WineRepository.GetWines(ctx, &wines)

	return wines, err
}
