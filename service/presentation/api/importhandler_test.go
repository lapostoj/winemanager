package api_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/presentation/api"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
	"github.com/lapostoj/winemanager/service/test"
)

type MockCsvImport struct {
	mock.Mock
}

func (mock *MockCsvImport) ExecuteCsvImport(ctx context.Context, reader *bufio.Reader) ([]wine.Wine, error) {
	args := mock.Called(ctx, reader)
	return args.Get(0).([]wine.Wine), args.Error(1)
}

func TestPostImport(t *testing.T) {
	ctx := context.Background()
	csvImport := new(MockCsvImport)
	handler := api.ImportHandler{CsvImport: csvImport}
	recorder := httptest.NewRecorder()
	csvFile := openCsvTestFile()
	csvImport.On("ExecuteCsvImport", ctx, mock.Anything).Return([]wine.Wine{test.AWine(), test.AWine()}, nil)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	createFormCsvFile(writer, "file", csvFile)
	writer.Close()

	request := httptest.NewRequest("POST", "/api/import", &body).WithContext(ctx)
	request.Header.Set("Origin", api.GetClientURL())
	request.Header.Set("Content-Type", writer.FormDataContentType())

	handler.PostImport(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	wineResponses := []response.WineResponse{}
	json.Unmarshal(buf.Bytes(), &wineResponses)

	assert.Equal(t, result.StatusCode, 201)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.GetClientURL())
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, len(wineResponses), 2)
}

func openCsvTestFile() *os.File {
	file, err := os.Open("../../test/test.csv")
	if err != nil {
		panic(err)
	}
	return file
}

func createFormCsvFile(w *multipart.Writer, fieldname string, file *os.File) (io.Writer, error) {
	quoteEscaper := strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			quoteEscaper.Replace(fieldname), quoteEscaper.Replace(file.Name())))
	h.Set("Content-Type", "text/csv")

	ioWriter, err := w.CreatePart(h)
	if err == nil {
		_, err = io.Copy(ioWriter, file)
	}

	return ioWriter, err
}
