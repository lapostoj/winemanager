package persistence

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// wineEntityKind defines the name of the entity in database.
const wineEntityKind = "Wine"

// WineRepository implements a Repository for the Wine table.
type WineRepository struct {
}

// GetWines returns all the wines in the table.
func (repository WineRepository) GetWines(ctx context.Context, wines *[]wine.Wine) error {
	q := datastore.NewQuery(wineEntityKind)

	_, err := DatastoreClient(ctx).GetAll(ctx, q, wines)
	return err
}

// SaveWine save the wine in the database.
func (repository WineRepository) SaveWine(ctx context.Context, wine *wine.Wine) (string, error) {
	key, err := DatastoreClient(ctx).Put(ctx, datastore.IncompleteKey(wineEntityKind, nil), wine)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}
