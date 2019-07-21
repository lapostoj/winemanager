package createcellar_test

import (
	"context"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/createcellar"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/lapostoj/winemanager/service/test"
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

func TestCreateCellar(t *testing.T) {
	ctx := context.Background()
	cellarRepository := new(MockCellarRepository)
	cellar := test.ACellar()
	expectedId := "123"
	createCellarService := createcellar.CreateCellar{CellarRepository: cellarRepository}
	cellarRepository.On("SaveCellar", ctx, &cellar).Return(expectedId, nil)

	id, err := createCellarService.Execute(ctx, &cellar)

	assert.Nil(t, err)
	assert.Equal(t, id, expectedId)
	cellarRepository.AssertCalled(t, "SaveCellar", ctx, &cellar)
}
