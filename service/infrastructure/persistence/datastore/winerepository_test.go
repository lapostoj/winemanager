package persistence_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/stretchr/testify/assert"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"github.com/lapostoj/winemanager/service/test"
)

func TestSaveWine(t *testing.T) {
	t.Skip("need to figure how to use the mock datastore")
	repository := persistence.WineRepository{}

	timeout, _ := time.ParseDuration("500ms")
	ctx, cancelFun := context.WithTimeout(nil, timeout)
	client, _ := datastore.NewClient(ctx, "my-project-id")
	aWine := test.AWine()

	encodedKey, err := repository.SaveWine(ctx, &aWine)
	assert.Nil(t, err)

	key, _ := datastore.DecodeKey(encodedKey)
	var retrievedWine wine.Wine
	client.Get(ctx, key, retrievedWine)

	assert.Equal(t, retrievedWine.Name, aWine.Name)
	cancelFun()
}
