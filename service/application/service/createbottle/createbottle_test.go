package createbottle_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/createbottle"
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
	return args.Error(0)
}

func TestCreateBottle(t *testing.T) {
	ctx := context.Background()
	bottleRepository := new(MockBottleRepository)
	bottle := test.ABottle()
	expectedId := "123"
	createBottleService := createbottle.CreateBottle{BottleRepository: bottleRepository}
	bottleRepository.On("SaveBottle", ctx, &bottle).Return(expectedId, nil)

	id, err := createBottleService.Execute(ctx, &bottle)

	assert.Nil(t, err)
	assert.Equal(t, id, expectedId)
	bottleRepository.AssertCalled(t, "SaveBottle", ctx, &bottle)
}
