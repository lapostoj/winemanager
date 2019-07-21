package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/lapostoj/winemanager/service/presentation/api"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/lapostoj/winemanager/service/test"
)

type MockGetCellar struct {
	mock.Mock
}

func (mock MockGetCellar) ForAccountID(ctx context.Context, accountID int) ([]cellar.Cellar, error) {
	args := mock.Called(ctx, accountID)
	return args.Get(0).([]cellar.Cellar), args.Error(1)
}

func TestQueryCellars(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getCellar := new(MockGetCellar)
	handler := api.CellarHandler{GetCellar: getCellar}

	var body bytes.Buffer
	request := httptest.NewRequest("GET", "/api/cellars?accountID=123", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientUrl())

	cellars := []cellar.Cellar{test.ACellar()}
	getCellar.On("ForAccountID", ctx, 123).Return(cellars, nil)

	handler.QueryCellars(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	cellarResponses := []response.CellarResponse{}
	json.Unmarshal(buf.Bytes(), &cellarResponses)

	assert.Equal(t, result.StatusCode, 200)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientUrl())
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, len(cellarResponses), 1)
}
