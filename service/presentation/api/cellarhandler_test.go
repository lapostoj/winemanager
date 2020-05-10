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

type MockCreateCellar struct {
	mock.Mock
}

func (mock MockGetCellar) ForAccountID(ctx context.Context, accountID int) ([]cellar.Cellar, error) {
	args := mock.Called(ctx, accountID)
	return args.Get(0).([]cellar.Cellar), args.Error(1)
}

func (mock MockCreateCellar) Execute(ctx context.Context, cellar *cellar.Cellar) (int64, error) {
	args := mock.Called(ctx, cellar)
	return args.Get(0).(int64), args.Error(1)
}

func TestQueryCellars(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getCellar := new(MockGetCellar)
	createCellar := new(MockCreateCellar)
	handler := api.CellarHandler{GetCellar: getCellar, CreateCellar: createCellar}

	var body bytes.Buffer
	request := httptest.NewRequest("GET", "/api/cellars?accountID=123", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	cellars := []cellar.Cellar{test.ACellarWithID()}
	getCellar.On("ForAccountID", ctx, 123).Return(cellars, nil)

	handler.QueryCellars(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	cellarResponses := []response.CellarResponse{}
	json.Unmarshal(buf.Bytes(), &cellarResponses)

	assert.Equal(t, result.StatusCode, 200)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientURL())
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, len(cellarResponses), 1)
}

func TestPostCellar(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getCellar := new(MockGetCellar)
	createCellar := new(MockCreateCellar)
	handler := api.CellarHandler{GetCellar: getCellar, CreateCellar: createCellar}
	expectedCellar := test.ACellar()
	ID := int64(1234)

	bodyBytes, err := json.Marshal(expectedCellar)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(bodyBytes)
	request := httptest.NewRequest("POST", "/api/cellars", body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	createCellar.On("Execute", ctx, mock.MatchedBy(func(cellar *cellar.Cellar) bool {
		return cellar.Name == expectedCellar.Name && cellar.AccountID == expectedCellar.AccountID
	})).Return(ID, nil)

	handler.PostCellar(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	IDResponse := response.IDResponse{}
	json.Unmarshal(buf.Bytes(), &IDResponse)

	assert.Equal(t, result.StatusCode, 201)
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, IDResponse.ID, ID)
}
