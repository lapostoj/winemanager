package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

// WineHandler defines the interface for a WineHandler
type WineHandler struct {
	WineRepository wine.Repository
}

// Test handles the GET calls to '/api/test'
func (handler WineHandler) Test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := handler.WineRepository.SaveTestWine(ctx); err != nil {
		log.Printf("TestWine - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("TestWine added"))
}

// GetWines handles the GET calls to '/api/wines' and return the known wines
func (handler WineHandler) GetWines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", GetClientURL())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var wines []wine.Wine
	if err := handler.WineRepository.GetWines(ctx, &wines); err != nil {
		log.Printf("GetWines - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(response.NewWineResponses(wines))
	if err != nil {
		log.Printf("GetWines - marshal: %q\n", err)
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

// PostWines handles the POST calls to '/api/wines' and add the wine in the database
func (handler WineHandler) PostWines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var postWineRequest request.PostWineRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postWineRequest)
	if err != nil {
		log.Printf("PostWines - decode: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wine := postWineRequest.NewWine()
	key, err := handler.WineRepository.SaveWine(ctx, wine)
	if err != nil {
		log.Printf("PostWines - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(key))
}
