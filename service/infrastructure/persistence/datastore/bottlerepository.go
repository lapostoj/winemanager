package persistence

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/lapostoj/winemanager/service/domain/model/bottle"
)

// bottleEntityKind defines the name of the entity in database.
const bottleEntityKind = "Bottle"

// BottleRepository implements a Repository for the Bottle table.
type BottleRepository struct {
}

// SaveBottle saves the bottle in the database
func (repository BottleRepository) SaveBottle(ctx context.Context, bottle *bottle.Bottle) (string, error) {
	key, err := DatastoreClient(ctx).Put(ctx, datastore.IncompleteKey(bottleEntityKind, nil), bottle)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}

// FindBottlesForCellarID returns the bottles in the table with the cellarID provided
func (repository BottleRepository) FindBottlesForCellarID(ctx context.Context, bottles *[]bottle.Bottle, cellarID int) error {
	q := datastore.NewQuery(bottleEntityKind).
		Filter("CellarID =", cellarID)

	_, err := DatastoreClient(ctx).GetAll(ctx, q, bottles)
	return err
}
