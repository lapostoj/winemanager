package response_test

import (
	"testing"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/stretchr/testify/assert"
)

func TestNewIDResponse(t *testing.T) {
	id := "id"

	IDResponse := response.NewIDResponse(id)

	assert.Equal(t, IDResponse.ID, id)
}
