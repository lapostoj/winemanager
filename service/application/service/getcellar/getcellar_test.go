package getcellar_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/getcellar"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCellarRepository struct {
	mock.Mock
}

func (mock *MockCellarRepository) SaveCellar(ctx context.Context, cellar *cellar.Cellar) (string, error) {
	args := mock.Called(ctx, cellar)
	return args.String(0), args.Error(1)
}

func (mock *MockCellarRepository) FindCellarsForAccountID(ctx context.Context, cellars *[]cellar.Cellar, accountId int) error {
	args := mock.Called(ctx, cellars, accountId)
	return args.Error(0)
}

func TestGetCellar(t *testing.T) {
	ctx := context.Background()
	cellarRepository := new(MockCellarRepository)
	accountId := 123
	var cellars []cellar.Cellar
	getCellarService := getcellar.GetCellar{CellarRepository: cellarRepository}
	cellarRepository.On("FindCellarsForAccountID", ctx, &cellars, accountId).Return(nil)

	cellars, err := getCellarService.ForAccountID(ctx, accountId)

	assert.Nil(t, err)
	cellarRepository.AssertCalled(t, "FindCellarsForAccountID", ctx, &cellars, accountId)
}
