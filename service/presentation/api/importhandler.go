package api

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/lapostoj/winemanager/service/application/service"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

type ImportHandler struct {
	CsvImport service.CsvImportInterface
}

// PostImport handles the POST calls to '/api/import' and parse the file to put it in the db.
// Inspired from https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html.
func (handler ImportHandler) PostImport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", GetClientUrl())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := validateHeaders(r); err != nil {
		log.Printf("PostImport - Invalid headers: %s", err.Error())
		http.Error(w, "Invalid headers", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, headers, err := r.FormFile("file")
	log.Printf("PostImport - Importing file: %s", headers.Filename)
	if err != nil {
		log.Printf("File: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()
	if err := validateFileType(headers); err != nil {
		log.Printf("PostImport - Invalid file: %s", err.Error())
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	wines, err := handler.CsvImport.ExecuteCsvImport(ctx, bufio.NewReader(file))
	if err != nil {
		log.Printf("PostImport - %s", err.Error())
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	log.Printf("Wines length - %d", len(wines))

	response, err := json.Marshal(response.NewWineResponses(wines))
	if err != nil {
		log.Printf("GetWines - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(response)
}

func validateHeaders(r *http.Request) error {
	originHeader := "Origin"
	if r.Header.Get(originHeader) == GetClientUrl() {
		return nil
	}
	return errors.New("'Origin': " + r.Header.Get(originHeader))
}

func validateFileType(fileHeaders *multipart.FileHeader) error {
	contentTypeHeader := "Content-Type"
	csvType := "text/csv"

	if fileHeaders.Header.Get(contentTypeHeader) == csvType {
		return nil
	}
	return errors.New("'Type': " + fileHeaders.Header.Get(contentTypeHeader))
}
