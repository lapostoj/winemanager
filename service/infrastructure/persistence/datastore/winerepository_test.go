package persistence_test

import (
	"testing"

	"google.golang.org/appengine/datastore"

	"github.com/stretchr/testify/assert"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"github.com/lapostoj/winemanager/service/test"
)

func TestSaveWine(t *testing.T) {
	t.Skip("persistence.SaveWine fails with 'service bridge HTTP failed'")

	ctx := test.AnAppEngineTestContext()
	aWine := test.AWine()

	encodedKey, err := persistence.SaveWine(ctx, &aWine)
	assert.Nil(t, err)

	key, _ := datastore.DecodeKey(encodedKey)
	var retrievedWine wine.Wine
	datastore.Get(ctx, key, retrievedWine)

	assert.Equal(t, retrievedWine.Name, aWine.Name)
}
