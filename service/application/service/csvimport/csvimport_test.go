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

type MockCellarRepository struct {
	mock.Mock
}

func (mock *MockCellarRepository) SaveCellar(ctx context.Context, cellar *cellar.Cellar) (string, error) {
	args := mock.Called(ctx, cellar)
	return args.String(0), args.Error(1)
}

func (mock *MockCellarRepository) FindCellarsForAccountID(ctx context.Context, cellars *[]cellar.Cellar, accountId int) error {
	args := mock.Called(ctx, cellars, accountId)
	*cellars = append(*cellars, test.ACellar(), test.ACellar())
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
	return args.Error(0)
}

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

func TestExecute(t *testing.T) {
	ctx := context.Background()
	cellarRepository := new(MockCellarRepository)
	wineRepository := new(MockWineRepository)
	bottleRepository := new(MockBottleRepository)
	csvImportService := csvimport.CsvImport{CellarRepository: cellarRepository, WineRepository: wineRepository, BottleRepository: bottleRepository}
	file := test.ACsvImportFile()
	reader := bufio.NewReader(strings.NewReader(file))
	cellarRepository.On("SaveCellar", ctx, mock.Anything).Return("id", nil)
	wineRepository.On("SaveWine", ctx, mock.Anything).Return("id", nil)
	bottleRepository.On("SaveBottle", ctx, mock.Anything).Return("id", nil)

	wines, err := csvImportService.Execute(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 3)
	cellarRepository.AssertCalled(t, "SaveCellar", ctx, mock.Anything)
	wineRepository.AssertCalled(t, "SaveWine", ctx, mock.Anything)
	bottleRepository.AssertCalled(t, "SaveBottle", ctx, mock.Anything)
}

func TestExecuteWithEmptyFile(t *testing.T) {
	ctx := context.Background()
	cellarRepository := new(MockCellarRepository)
	wineRepository := new(MockWineRepository)
	bottleRepository := new(MockBottleRepository)
	csvImportService := csvimport.CsvImport{CellarRepository: cellarRepository, WineRepository: wineRepository, BottleRepository: bottleRepository}
	file := ""
	reader := bufio.NewReader(strings.NewReader(file))

	wines, err := csvImportService.Execute(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 0)
	cellarRepository.AssertNotCalled(t, "SaveCellar")
	wineRepository.AssertNotCalled(t, "SaveWine")
	bottleRepository.AssertNotCalled(t, "SaveBottle")
}
