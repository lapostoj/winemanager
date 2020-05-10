package getcellar_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/getcellar"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/lapostoj/winemanager/service/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCellarRepository struct {
	mock.Mock
}

func (mock *MockCellarRepository) SaveCellar(ctx context.Context, cellar *cellar.Cellar) (int64, error) {
	args := mock.Called(ctx, cellar)
	return args.Get(0).(int64), args.Error(1)
}

func (mock *MockCellarRepository) FindCellarsForAccountID(ctx context.Context, cellars *[]cellar.Cellar, accountId int) error {
	args := mock.Called(ctx, cellars, accountId)
	*cellars = append(*cellars, test.ACellar(), test.ACellar())
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
	assert.Equal(t, len(cellars), 2)
	cellarRepository.AssertCalled(t, "FindCellarsForAccountID", ctx, &cellars, accountId)
}
