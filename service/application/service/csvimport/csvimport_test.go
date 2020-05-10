package csvimport_test

import (
	"bufio"
	"context"
	"strings"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/csvimport"
	"github.com/lapostoj/winemanager/service/domain/model/bottle"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateCellar struct {
	mock.Mock
}

func (mock *MockCreateCellar) Execute(ctx context.Context, cellar *cellar.Cellar) (int64, error) {
	args := mock.Called(ctx, cellar)
	return args.Get(0).(int64), args.Error(1)
}

type MockCreateWine struct {
	mock.Mock
}

func (mock *MockCreateWine) Execute(ctx context.Context, wine *wine.Wine) (int64, error) {
	args := mock.Called(ctx, wine)
	return args.Get(0).(int64), args.Error(1)
}

type MockCreateBottle struct {
	mock.Mock
}

func (mock *MockCreateBottle) Execute(ctx context.Context, bottle *bottle.Bottle) (int64, error) {
	args := mock.Called(ctx, bottle)
	return args.Get(0).(int64), args.Error(1)
}

func TestExecute(t *testing.T) {
	ctx := context.Background()
	createCellar := new(MockCreateCellar)
	createWine := new(MockCreateWine)
	createBottle := new(MockCreateBottle)
	csvImportService := csvimport.CsvImport{CreateCellar: createCellar, CreateWine: createWine, CreateBottle: createBottle}
	file := test.ACsvImportFile()
	reader := bufio.NewReader(strings.NewReader(file))
	createCellar.On("Execute", ctx, mock.Anything).Return(int64(111), nil)
	createWine.On("Execute", ctx, mock.Anything).Return(int64(122), nil)
	createBottle.On("Execute", ctx, mock.Anything).Return(int64(133), nil)

	wines, err := csvImportService.Execute(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 3)
	createCellar.AssertNumberOfCalls(t, "Execute", 1)
	createCellar.AssertCalled(t, "Execute", ctx, mock.Anything)
	createWine.AssertNumberOfCalls(t, "Execute", 3)
	createWine.AssertCalled(t, "Execute", ctx, mock.Anything)
	createBottle.AssertNumberOfCalls(t, "Execute", 3)
	createBottle.AssertCalled(t, "Execute", ctx, mock.Anything)
}

func TestExecuteWithEmptyFile(t *testing.T) {
	ctx := context.Background()
	createCellar := new(MockCreateCellar)
	createWine := new(MockCreateWine)
	createBottle := new(MockCreateBottle)
	csvImportService := csvimport.CsvImport{CreateCellar: createCellar, CreateWine: createWine, CreateBottle: createBottle}
	file := ""
	reader := bufio.NewReader(strings.NewReader(file))

	wines, err := csvImportService.Execute(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 0)
	createCellar.AssertNotCalled(t, "Execute")
	createWine.AssertNotCalled(t, "Execute")
	createBottle.AssertNotCalled(t, "Execute")
}
