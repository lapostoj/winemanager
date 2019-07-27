package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/lapostoj/winemanager/service/application/service/createbottle"
	"github.com/lapostoj/winemanager/service/application/service/getbottles"
	"github.com/lapostoj/winemanager/service/infrastructure/utils"

	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

// BottleHandlerInterface defines the interface for a BottleHandler
type BottleHandlerInterface interface {
	QueryBottles(w http.ResponseWriter, r *http.Request)
	PostBottle(w http.ResponseWriter, r *http.Request)
}

// BottleHandler implements handling of api calls for bottles
type BottleHandler struct {
	GetBottles   getbottles.GetBottlesService
	CreateBottle createbottle.CreateBottleService
}

// QueryBottles handles the GET calls to '/api/bottles' and return the bottles matching the query
func (handler BottleHandler) QueryBottles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", GetClientURL())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	values, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil || values.Get("cellarID") == "" {
		log.Printf("QueryBottles - parseQuery: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cellarID := utils.StringToInt(values.Get("cellarID"))

	bottles, err := handler.GetBottles.ForCellarID(ctx, cellarID)
	if err != nil {
		log.Printf("QueryBottles - service: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(response.NewBottleResponses(bottles))
	if err != nil {
		log.Printf("QueryBottles - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

// PostBottle handles the POST calls to '/api/bottles' and add the bottle in the database
func (handler BottleHandler) PostBottle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", GetClientURL())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var postBottleRequest request.PostBottleRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postBottleRequest)
	if err != nil {
		log.Printf("PostBottle - decode: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bottle := postBottleRequest.NewBottle()
	key, err := handler.CreateBottle.Execute(ctx, bottle)
	if err != nil {
		log.Printf("PostBottle - service: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(response.NewIDResponse(key))
	if err != nil {
		log.Printf("PostBottle - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(response)
}
