package persistence

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// entityKind defines the name of the entity in database.
const entityKind = "Wine"

// WineRepository implements a Repository for the Wine table.
type WineRepository struct {
}

// SaveTestWine adds a test entry in the Wine table.
func (repository WineRepository) SaveTestWine(ctx context.Context) error {
	testWine := wine.Wine{
		Name:        "Test Wine",
		Designation: "Test Designation",
		Growth:      "Test Growth",
		Country:     "FR",
		Region:      "Bourgogne",
		Producer:    "Test Producer",
		Color:       wine.RED,
		Type:        wine.SEC,
	}

	_, err := repository.SaveWine(ctx, &testWine)
	return err
}

// GetWinesInStock returns the wines in the table with a non 0 quantity.
func (repository WineRepository) GetWinesInStock(ctx context.Context, wines *[]wine.Wine) error {
	q := datastore.NewQuery(entityKind).Filter("Quantity >", 0)

	_, err := DatastoreClient(ctx).GetAll(ctx, q, wines)
	return err
}

// SaveWine save the wine in the database.
func (repository WineRepository) SaveWine(ctx context.Context, wine *wine.Wine) (string, error) {
	key, err := DatastoreClient(ctx).Put(ctx, datastore.IncompleteKey(entityKind, nil), wine)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}
