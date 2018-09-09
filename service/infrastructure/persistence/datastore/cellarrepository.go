package persistence

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
)

// CellarRepository implements a Repository for the Cellar table.
type CellarRepository struct {
}

// SaveCellar saves the cellar in the database
func (repository CellarRepository) SaveCellar(ctx context.Context, cellar *cellar.Cellar) (string, error) {
	return "", nil
}

// FindCellarsForAccountID returns the cellars in the table with the accountID provided
func (repository CellarRepository) FindCellarsForAccountID(ctx context.Context, cellars *[]cellar.Cellar, accountID int) error {
	q := datastore.NewQuery(entityKind).
		Filter("AccountID =", accountID)

	_, err := DatastoreClient(ctx).GetAll(ctx, q, cellars)
	return err
}
