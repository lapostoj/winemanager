package api

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/lapostoj/winemanager/service/application/service"
	"github.com/lapostoj/winemanager/service/presentation/api/response"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// PostImport handles the POST calls to '/api/import' and parse the file to put it in the db.
// Inspired from https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html.
func PostImport(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	w.Header().Set("Access-Control-Allow-Origin", website)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := validateHeaders(r); err != nil {
		log.Warningf(ctx, "PostImport - Invalid headers: %s", err.Error())
		http.Error(w, "Invalid headers", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, headers, err := r.FormFile("file")
	filename := headers.Filename
	log.Infof(ctx, "PostImport - Importing file: %s", filename)
	if err != nil {
		log.Errorf(ctx, "File: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()
	if err := validateFileType(headers); err != nil {
		log.Warningf(ctx, "PostImport - Invalid file: %s", err.Error())
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	wines, err := service.ExecuteCsvImport(ctx, file)
	if err != nil {
		log.Warningf(ctx, "PostImport - %s", err.Error())
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(response.NewWineResponses(wines))
	if err != nil {
		log.Errorf(ctx, "GetWines - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(response)
}

func validateHeaders(r *http.Request) error {
	originHeader := "Origin"
	if r.Header.Get(originHeader) == website {
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
