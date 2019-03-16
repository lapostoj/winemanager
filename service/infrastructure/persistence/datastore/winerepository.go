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

// TestWine adds a test entry in the Wine table.
func TestWine(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, "my-project-id")
	if err != nil {
		log.Println("Datastore client error")
		log.Fatal(err)
	}

	k := datastore.IncompleteKey(entityKind, nil)

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

	_, err2 := client.Put(ctx, k, testWine)
	return err2
}

// GetWinesInStock returns the wines in the table with a non 0 quantity.
func GetWinesInStock(ctx context.Context, wines *[]wine.Wine) error {
	client, err := datastore.NewClient(ctx, "my-project-id")
	if err != nil {
		log.Fatal(err)
	}

	q := datastore.NewQuery(entityKind).Filter("Quantity >", 0)

	_, err2 := client.GetAll(ctx, q, wines)
	return err2
}

// SaveWine save the wine in the database.
func SaveWine(ctx context.Context, wine *wine.Wine) (string, error) {
	client, err := datastore.NewClient(ctx, "my-project-id")
	if err != nil {
		log.Fatal(err)
	}

	key := datastore.IncompleteKey(entityKind, nil)
	if _, err2 := client.Put(ctx, key, wine); err != nil {
		return "", err2
	}
	return key.Encode(), nil
}
