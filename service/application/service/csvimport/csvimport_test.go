package csvimport_test

import (
	"bufio"
	"context"
	"strings"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service/csvimport"
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

func TestExecute(t *testing.T) {
	ctx := context.Background()
	wineRepository := new(MockWineRepository)
	csvImportService := csvimport.CsvImport{WineRepository: wineRepository}
	file := test.ACsvImportFile()
	reader := bufio.NewReader(strings.NewReader(file))
	wineRepository.On("SaveWine", ctx, mock.Anything).Return("id", nil)

	wines, err := csvImportService.Execute(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 3)
	wineRepository.AssertCalled(t, "SaveWine", ctx, mock.Anything)
}

func TestExecuteWithEmptyFile(t *testing.T) {
	ctx := context.Background()
	wineRepository := new(MockWineRepository)
	csvImportService := csvimport.CsvImport{WineRepository: wineRepository}
	file := ""
	reader := bufio.NewReader(strings.NewReader(file))

	wines, err := csvImportService.Execute(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 0)
	wineRepository.AssertNotCalled(t, "SaveWine")
}
