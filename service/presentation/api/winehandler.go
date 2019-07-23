package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lapostoj/winemanager/service/application/service/createwine"
	"github.com/lapostoj/winemanager/service/application/service/getwines"
	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

// WineHandlerInterface defines the interface for a WineHandler
type WineHandlerInterface interface {
	QueryWines(w http.ResponseWriter, r *http.Request)
	OptionsWines(w http.ResponseWriter, r *http.Request)
	PostWine(w http.ResponseWriter, r *http.Request)
	PostTest(w http.ResponseWriter, r *http.Request)
}

// WineHandler defines the interface for a WineHandler
type WineHandler struct {
	// WineRepository wine.Repository
	GetWines   getwines.GetWinesService
	CreateWine createwine.CreateWineService
}

// QueryWines handles the GET calls to '/api/wines' and return the known wines
func (handler WineHandler) QueryWines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", GetClientURL())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	wines, err := handler.GetWines.Execute(ctx)
	if err != nil {
		log.Printf("QueryWines - service: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(response.NewWineResponses(wines))
	if err != nil {
		log.Printf("QueryWines - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

// OptionsWines handles the OPTIONS calls to '/api/wines' and check their headers
func (handler WineHandler) OptionsWines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", GetClientURL())
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// PostWine handles the POST calls to '/api/wines' and add the wine in the database
func (handler WineHandler) PostWine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var postWineRequest request.PostWineRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postWineRequest)
	if err != nil {
		log.Printf("PostWine - decode: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wine := postWineRequest.NewWine()
	key, err := handler.CreateWine.Execute(ctx, wine)
	if err != nil {
		log.Printf("PostWine - service: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(key))
}

// PostTest handles the GET calls to '/api/test'
func (handler WineHandler) PostTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	testWine := wine.Wine{
		Name:        "Test Wine",
		Designation: "Test Designation",
		Growth:      "Test Growth",
		Country:     "FR",
		Region:      "Bourgogne",
		Producer:    "Test Producer",
		Color:       wine.RED,
		Type:        wine.SEC,
	}
	key, err := handler.CreateWine.Execute(ctx, &testWine)
	if err != nil {
		log.Printf("PostTest - service: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(key))
}
