package getbottles_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/getbottles"
	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/lapostoj/winemanager/service/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBottleRepository struct {
	mock.Mock
}

func (mock *MockBottleRepository) SaveBottle(ctx context.Context, bottle *bottle.Bottle) (string, error) {
	args := mock.Called(ctx, bottle)
	return args.String(0), args.Error(1)
}

func (mock *MockBottleRepository) FindBottlesForCellarID(ctx context.Context, bottles *[]bottle.Bottle, cellarId int) error {
	args := mock.Called(ctx, bottles, cellarId)
	*bottles = append(*bottles, test.ABottle(), test.ABottle())
	return args.Error(0)
}

type MockWineRepository struct {
	mock.Mock
}

func (mock *MockWineRepository) SaveWine(ctx context.Context, wine *wine.Wine) (string, error) {
	args := mock.Called(ctx, wine)
	return args.String(0), args.Error(1)
}

func (mock *MockWineRepository) GetWines(ctx context.Context, wines *[]wine.Wine) error {
	args := mock.Called(ctx, wines)
	return args.Error(0)
}

func (mock *MockWineRepository) GetWineByID(ctx context.Context, ID int64, wine *wine.Wine) error {
	args := mock.Called(ctx, ID, wine)
	*wine = test.AWineWithID()
	return args.Error(0)
}

func TestGetBottles(t *testing.T) {
	ctx := context.Background()
	bottleRepository := new(MockBottleRepository)
	wineRepository := new(MockWineRepository)
	cellarId := 123
	wineId := int64(111)
	var bottleResponses []response.BottleResponse
	getBottleService := getbottles.GetBottles{BottleRepository: bottleRepository, WineRespository: wineRepository}
	bottleRepository.On("FindBottlesForCellarID", ctx, mock.AnythingOfType("*[]bottle.Bottle"), cellarId).Return(nil)
	wineRepository.On("GetWineByID", ctx, wineId, mock.AnythingOfType("*wine.Wine")).Return(nil)

	bottleResponses, err := getBottleService.ForCellarID(ctx, cellarId)

	assert.Nil(t, err)
	assert.Equal(t, len(bottleResponses), 2)
	bottleRepository.AssertCalled(t, "FindBottlesForCellarID", ctx, mock.AnythingOfType("*[]bottle.Bottle"), cellarId)
	wineRepository.AssertCalled(t, "GetWineByID", ctx, wineId, mock.AnythingOfType("*wine.Wine"))
}
