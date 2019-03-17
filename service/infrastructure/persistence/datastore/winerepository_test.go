package persistence_test

import (
	"context"
	"log"
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
	if err := datastoreClient(ctx).Get(ctx, key, &retrievedWine); err != nil {
		panic(err)
	}

	assert.Equal(t, retrievedWine.Name, aWine.Name)
	assert.Equal(t, retrievedWine.CreationTime.UTC().String(), aWine.CreationTime.UTC().String())
}

func datastoreClient(ctx context.Context) *datastore.Client {
	client, err := datastore.NewClient(ctx, "cave-inventaire")
	if err != nil {
		log.Println("Datastore client error")
		log.Fatal(err)
	}
	return client
}
