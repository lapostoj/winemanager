package request_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/stretchr/testify/assert"
)

func TestNewCellar(t *testing.T) {
	request := request.PostCellarRequest{
		Name:      "Cellar",
		AccountID: 111,
	}

	cellar := request.NewCellar()

	assert.Equal(t, request.Name, cellar.Name)
	assert.Equal(t, request.AccountID, cellar.AccountID)
	assert.NotNil(t, cellar.CreationTime)
}
