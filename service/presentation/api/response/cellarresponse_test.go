package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/lapostoj/winemanager/service/test"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewCellarResponse(t *testing.T) {
	cellar := test.ACellarWithID()

	cellarResponse := response.NewCellarResponse(cellar)

	assert.Equal(t, cellarResponse.ID, cellar.Key.ID)
	assert.Equal(t, cellarResponse.Name, cellar.Name)
	assert.Equal(t, cellarResponse.AccountID, cellar.AccountID)
	assert.Equal(t, cellarResponse.CreationTime, cellar.CreationTime.String())
}

func TestNewCellarResponses(t *testing.T) {
	cellars := []cellar.Cellar{test.ACellarWithID()}

	cellarResponses := response.NewCellarResponses(cellars)

	assert.Equal(t, len(cellarResponses), 1)
	assert.Equal(t, cellarResponses[0].ID, cellars[0].Key.ID)
	assert.Equal(t, cellarResponses[0].Name, cellars[0].Name)
	assert.Equal(t, cellarResponses[0].AccountID, cellars[0].AccountID)
	assert.Equal(t, cellarResponses[0].CreationTime, cellars[0].CreationTime.String())
}
