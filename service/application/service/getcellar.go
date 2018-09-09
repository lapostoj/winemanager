package service

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/cellar"
)

// GetCellarService defines the interface for a GetCellarService
type GetCellarService interface {
	ForAccountID(ctx context.Context, accountID int) ([]cellar.Cellar, error)
}

// GetCellar implements a service to get Cellars
type GetCellar struct {
	Repository cellar.Repository
}

// ForAccountID returns the cellars with the provided accountID.
func (service GetCellar) ForAccountID(ctx context.Context, accountID int) ([]cellar.Cellar, error) {
	var cellars []cellar.Cellar
	err := service.Repository.FindCellarsForAccountID(ctx, &cellars, accountID)

	return cellars, err
}
