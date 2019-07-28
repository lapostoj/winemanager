package createwine_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/createwine"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	return args.Error(0)
}

func TestCreateCellar(t *testing.T) {
	ctx := context.Background()
	wineRepository := new(MockWineRepository)
	wine := test.AWine()
	expectedId := "123"
	createWineService := createwine.CreateWine{WineRepository: wineRepository}
	wineRepository.On("SaveWine", ctx, &wine).Return(expectedId, nil)

	id, err := createWineService.Execute(ctx, &wine)

	assert.Nil(t, err)
	assert.Equal(t, id, expectedId)
	wineRepository.AssertCalled(t, "SaveWine", ctx, &wine)
}
