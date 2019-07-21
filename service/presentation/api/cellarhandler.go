package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/lapostoj/winemanager/service/application/service/getcellar"
	"github.com/lapostoj/winemanager/service/application/service/createcellar"
	"github.com/lapostoj/winemanager/service/infrastructure/utils"

	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

// CellarHanderInterface defines the interface for a CellarHanderInterface
type CellarHanderInterface interface {
	QueryCellars(w http.ResponseWriter, r *http.Request)
	PostCellar(w http.ResponseWriter, r *http.Request)
}

// CellarHandler implements handling of api calls for cellars
type CellarHandler struct {
	GetCellar    getcellar.GetCellarService
	CreateCellar createcellar.CreateCellarService
}

// QueryCellars handles the GET calls to '/api/cellars' and return the cellars matching the query
func (handler CellarHandler) QueryCellars(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", GetClientURL())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	values, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil || values.Get("accountID") == "" {
		log.Printf("GetCellars - parseQuery: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountID := utils.StringToInt(values.Get("accountID"))

	cellars, err := handler.GetCellar.ForAccountID(ctx, accountID)
	if err != nil {
		log.Printf("GetCellars - service: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(response.NewCellarResponses(cellars))
	if err != nil {
		log.Printf("GetCellars - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

// PostCellar handles the POST calls to '/api/cellars' and add the cellar in the database
func (handler CellarHandler) PostCellar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var postCellarRequest request.PostCellarRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postCellarRequest)
	if err != nil {
		log.Printf("PostCellar - decode: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cellar := postCellarRequest.NewCellar()
	key, err := handler.CreateCellar.Execute(ctx, cellar)
	if err != nil {
		log.Printf("PostCellar - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(key))
}
