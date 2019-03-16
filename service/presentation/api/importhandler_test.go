package api_test

import (
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

	"github.com/lapostoj/winemanager/service/presentation/api"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

func TestPostImport(t *testing.T) {
	ctx := context.Background()
	recorder := httptest.NewRecorder()
	csvFile := openCsvTestFile()

	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	createFormCsvFile(w, "file", csvFile)
	w.Close()

	request := httptest.NewRequest("POST", "/api/import", &body).WithContext(ctx)
	request.Header.Set("Origin", api.Website)
	request.Header.Set("Content-Type", w.FormDataContentType())

	api.PostImport(recorder, request)

	buf := new(bytes.Buffer)
	result := recorder.Result()
	buf.ReadFrom(result.Body)
	wineResponses := []response.WineResponse{}
	json.Unmarshal(buf.Bytes(), &wineResponses)

	assert.Equal(t, result.StatusCode, 201)
	assert.Equal(t, result.Header.Get("Access-Control-Allow-Origin"), api.Website)
	assert.Equal(t, result.Header.Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, len(wineResponses), 3)
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
