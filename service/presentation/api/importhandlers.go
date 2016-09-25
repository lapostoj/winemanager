package api

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/lapostoj/winemanager/service/application/service"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// PostImport handles the POST calls to '/api/import' and parse the file to put it in the db.
// Inspired from https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html.
func PostImport(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if err := validateHeaders(r); err != nil {
		log.Warningf(ctx, "PostImport: Invalid headers: %s", err.Error())
		http.Error(w, "Invalid headers", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if err := validateFileType(handler); err != nil {
		log.Warningf(ctx, "PostImport: Invalid file: %s", err.Error())
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	if err := service.ExecuteCsvImport(ctx, file); err != nil {
		log.Warningf(ctx, "PostImport - %s", err.Error())
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(handler.Filename + " imported."))
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
	// Whatever I tried I never got 'text/csv' but always 'application/vnd.ms-excel'.
	// Can be changed later if some 'test/csv' appear.
	csvType := "application/vnd.ms-excel"

	if fileHeaders.Header.Get(contentTypeHeader) == csvType {
		return nil
	}
	return errors.New("'Type': " + fileHeaders.Header.Get(contentTypeHeader))
}
