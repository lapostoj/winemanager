package persistence

import (
	"context"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"google.golang.org/appengine/datastore"
)

// entityKind defines the name of the entity in database.
const entityKind = "Wine"

// WineRepository implements a Repository for the Wine table.
type WineRepository struct {
}

// TestWine adds a test entry in the Wine table.
func TestWine(ctx context.Context) error {
	k := datastore.NewIncompleteKey(ctx, entityKind, nil)

	storageLocation := wine.StorageLocation{Cellar: "Moiré", Position: "3"}
	testWine := wine.NewWine()
	testWine.Name = "Test Wine"
	testWine.Designation = "Test Designation"
	testWine.Growth = "Test Growth"
	testWine.Country = "FR"
	testWine.Region = "Bourgogne"
	testWine.Producer = "Test Producer"
	testWine.Color = wine.RED
	testWine.Type = wine.SEC
	// Following should be on a "Bottle" object
	// We can have several bottle of the same type of wines.
	testWine.Year = 1963
	testWine.Quantity = 3
	testWine.Size = wine.BOTTLE
	testWine.StorageLocation = storageLocation

	_, err := datastore.Put(ctx, k, testWine)
	return err
}

// GetWinesInStock returns the wines in the table with a non 0 quantity.
func GetWinesInStock(ctx context.Context, wines *[]wine.Wine) error {
	q := datastore.NewQuery(entityKind).
		Filter("Quantity >", 0)

	_, err := q.GetAll(ctx, wines)
	return err
}

// SaveWine save the wine in the database.
func SaveWine(ctx context.Context, wine *wine.Wine) (string, error) {
	key := datastore.NewIncompleteKey(ctx, entityKind, nil)
	if _, err := datastore.Put(ctx, key, wine); err != nil {
		return "", err
	}
	return key.Encode(), nil
}
