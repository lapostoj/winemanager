package persistence_test

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"

	"github.com/stretchr/testify/assert"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"github.com/lapostoj/winemanager/service/test"
)

func TestSaveWine(t *testing.T) {
	t.Skip("Need to run the datastore emulator to run it")

	repository := persistence.WineRepository{}
	aWine := test.AWine()
	ctx := context.Background()

	encodedKey, err := repository.SaveWine(ctx, &aWine)
	assert.Nil(t, err)

	key, _ := datastore.DecodeKey(encodedKey)
	var retrievedWine wine.Wine
	if err := persistence.DatastoreClient(ctx).Get(ctx, key, &retrievedWine); err != nil {
		panic(err)
	}

	assert.Equal(t, retrievedWine, aWine)
}
