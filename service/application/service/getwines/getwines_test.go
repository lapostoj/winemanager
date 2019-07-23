package getwines_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/getwines"
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
	*wines = append(*wines, test.AWine(), test.AWine())
	return args.Error(0)
}

func TestGetWines(t *testing.T) {
	ctx := context.Background()
	wineRepository := new(MockWineRepository)
	var wines []wine.Wine
	getWinesService := getwines.GetWines{WineRepository: wineRepository}
	wineRepository.On("GetWines", ctx, &wines).Return(nil)

	wines, err := getWinesService.Execute(ctx)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 2)
	wineRepository.AssertCalled(t, "GetWines", ctx, &wines)
}
