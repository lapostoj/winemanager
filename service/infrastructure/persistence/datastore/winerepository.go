package persistence

import (
	"context"
	"log"

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

	_, err := repository.SaveWine(ctx, testWine)
	return err
}

// GetWinesInStock returns the wines in the table with a non 0 quantity.
func (repository WineRepository) GetWinesInStock(ctx context.Context, wines *[]wine.Wine) error {
	q := datastore.NewQuery(entityKind).Filter("Quantity >", 0)

	_, err := datastoreClient(ctx).GetAll(ctx, q, wines)
	return err
}

// SaveWine save the wine in the database.
func (repository WineRepository) SaveWine(ctx context.Context, wine *wine.Wine) (string, error) {
	key, err := datastoreClient(ctx).Put(ctx, datastore.IncompleteKey(entityKind, nil), wine)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}

func datastoreClient(ctx context.Context) *datastore.Client {
	client, err := datastore.NewClient(ctx, "cave-inventaire")
	if err != nil {
		log.Println("Datastore client error")
		log.Fatal(err)
	}
	return client
}
