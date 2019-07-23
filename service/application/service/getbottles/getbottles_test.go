package getbottles_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/getbottles"
	"github.com/lapostoj/winemanager/service/domain/model/bottle"
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

func TestGetBottles(t *testing.T) {
	ctx := context.Background()
	bottleRepository := new(MockBottleRepository)
	cellarId := 123
	var bottles []bottle.Bottle
	getBottleService := getbottles.GetBottles{BottleRepository: bottleRepository}
	bottleRepository.On("FindBottlesForCellarID", ctx, &bottles, cellarId).Return(nil)

	bottles, err := getBottleService.ForCellarID(ctx, cellarId)

	assert.Nil(t, err)
	assert.Equal(t, len(bottles), 2)
	bottleRepository.AssertCalled(t, "FindBottlesForCellarID", ctx, &bottles, cellarId)
}
