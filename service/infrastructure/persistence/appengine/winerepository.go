package persistence

import (
	"golang.org/x/net/context"

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
	testWine.StorageLocation = storageLocation
	testWine.Quantity = 3
	testWine.Color = wine.ROSE

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
