package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/presentation/api"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/lapostoj/winemanager/service/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetWines struct {
	mock.Mock
}

type MockCreateWine struct {
	mock.Mock
}

func (mock MockGetWines) Execute(ctx context.Context) ([]wine.Wine, error) {
	args := mock.Called(ctx)
	return args.Get(0).([]wine.Wine), args.Error(1)
}

func (mock MockCreateWine) Execute(ctx context.Context, wine *wine.Wine) (int64, error) {
	args := mock.Called(ctx, wine)
	return args.Get(0).(int64), args.Error(1)
}

func TestGetWines(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getWines := new(MockGetWines)
	createWine := new(MockCreateWine)
	handler := api.WineHandler{GetWines: getWines, CreateWine: createWine}

	var body bytes.Buffer
	request := httptest.NewRequest("GET", "/api/wines", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	wines := []wine.Wine{test.AWineWithID()}
	getWines.On("Execute", ctx).Return(wines, nil)

	handler.QueryWines(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	wineResponses := []response.WineResponse{}
	json.Unmarshal(buf.Bytes(), &wineResponses)

	assert.Equal(t, result.StatusCode, 200)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientURL())
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, len(wineResponses), 1)
}

func TestOptionWines(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getWines := new(MockGetWines)
	createWine := new(MockCreateWine)
	handler := api.WineHandler{GetWines: getWines, CreateWine: createWine}

	var body bytes.Buffer
	request := httptest.NewRequest("OPTIONS", "/api/wines", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	handler.OptionsWines(recorder, request)

	result := recorder.Result()

	assert.Equal(t, result.StatusCode, 200)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientURL())
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Methods"), "POST")
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Headers"), "Content-Type")
}

func TestPostWine(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getWines := new(MockGetWines)
	createWine := new(MockCreateWine)
	handler := api.WineHandler{GetWines: getWines, CreateWine: createWine}
	expectedWine := test.AWine()
	ID := int64(1234)

	bodyBytes, err := json.Marshal(expectedWine)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(bodyBytes)
	request := httptest.NewRequest("POST", "/api/wines", body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	createWine.On("Execute", ctx, &expectedWine).Return(ID, nil)

	handler.PostWine(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	IDResponse := response.IDResponse{}
	json.Unmarshal(buf.Bytes(), &IDResponse)

	assert.Equal(t, result.StatusCode, 201)
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, IDResponse.ID, ID)
}

func TestPostTest(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getWines := new(MockGetWines)
	createWine := new(MockCreateWine)
	handler := api.WineHandler{GetWines: getWines, CreateWine: createWine}
	expectedWine := test.AWine()
	ID := int64(1234)

	var body bytes.Buffer
	request := httptest.NewRequest("POST", "/api/test", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	createWine.On("Execute", ctx, &expectedWine).Return(ID, nil)

	handler.PostTest(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	IDResponse := response.IDResponse{}
	json.Unmarshal(buf.Bytes(), &IDResponse)

	assert.Equal(t, result.StatusCode, 201)
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, IDResponse.ID, ID)
}
