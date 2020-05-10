package persistence

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
)

// cellarEntityKind defines the name of the entity in database.
const cellarEntityKind = "Cellar"

// CellarRepository implements a Repository for the Cellar table.
type CellarRepository struct {
}

// SaveCellar saves the cellar in the database
func (repository CellarRepository) SaveCellar(ctx context.Context, cellar *cellar.Cellar) (int64, error) {
	key, err := DatastoreClient(ctx).Put(ctx, datastore.IncompleteKey(cellarEntityKind, nil), cellar)
	if err != nil {
		return -1, err
	}
	return key.ID, nil
}

// FindCellarsForAccountID returns the cellars in the table with the accountID provided
func (repository CellarRepository) FindCellarsForAccountID(ctx context.Context, cellars *[]cellar.Cellar, accountID int) error {
	q := datastore.NewQuery(cellarEntityKind).
		Filter("AccountID =", accountID)

	_, err := DatastoreClient(ctx).GetAll(ctx, q, cellars)
	return err
}
