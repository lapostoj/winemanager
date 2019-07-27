package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/presentation/api"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/lapostoj/winemanager/service/test"
)

type MockGetBottles struct {
	mock.Mock
}

type MockCreateBottle struct {
	mock.Mock
}

func (mock MockGetBottles) ForCellarID(ctx context.Context, cellarID int) ([]bottle.Bottle, error) {
	args := mock.Called(ctx, cellarID)
	return args.Get(0).([]bottle.Bottle), args.Error(1)
}

func (mock MockCreateBottle) Execute(ctx context.Context, bottle *bottle.Bottle) (string, error) {
	args := mock.Called(ctx, bottle)
	return args.Get(0).(string), args.Error(1)
}

func TestQueryBottles(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getBottles := new(MockGetBottles)
	createBottle := new(MockCreateBottle)
	handler := api.BottleHandler{GetBottles: getBottles, CreateBottle: createBottle}

	var body bytes.Buffer
	request := httptest.NewRequest("GET", "/api/bottles?cellarID=123", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	bottles := []bottle.Bottle{test.ABottle()}
	getBottles.On("ForCellarID", ctx, 123).Return(bottles, nil)

	handler.QueryBottles(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	bottleResponses := []response.BottleResponse{}
	json.Unmarshal(buf.Bytes(), &bottleResponses)

	assert.Equal(t, result.StatusCode, 200)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientURL())
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, len(bottleResponses), 1)
}

func TestPostBottle(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	getBottles := new(MockGetBottles)
	createBottle := new(MockCreateBottle)
	handler := api.BottleHandler{GetBottles: getBottles, CreateBottle: createBottle}
	expectedBottle := test.ABottle()
	ID := "id"

	bodyBytes, err := json.Marshal(expectedBottle)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(bodyBytes)
	request := httptest.NewRequest("POST", "/api/bottles", body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())

	createBottle.On("Execute", ctx, mock.MatchedBy(func(bottle *bottle.Bottle) bool {
		return bottle.Year == expectedBottle.Year && bottle.CellarID == expectedBottle.CellarID && bottle.Quantity == expectedBottle.Quantity
	})).Return(ID, nil)

	handler.PostBottle(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	IDResponse := response.IDResponse{}
	json.Unmarshal(buf.Bytes(), &IDResponse)

	assert.Equal(t, result.StatusCode, 201)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientURL())
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, IDResponse.ID, ID)
}
