package service_test

import (
	"bufio"
	"context"
	"strings"
	"testing"

	"github.com/lapostoj/winemanager/service/application/service"
	"github.com/lapostoj/winemanager/service/test"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCsvImport(t *testing.T) {
	file := test.ACsvImportFile()
	reader := bufio.NewReader(strings.NewReader(file))
	ctx := context.Background()

	wines, err := service.ExecuteCsvImport(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 3)
}

func TestExecuteCsvImportEmptyFile(t *testing.T) {
	file := ""
	reader := bufio.NewReader(strings.NewReader(file))
	ctx := context.Background()

	wines, err := service.ExecuteCsvImport(ctx, reader)

	assert.Nil(t, err)
	assert.Equal(t, len(wines), 0)
}
