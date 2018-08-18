package api

import (
	"encoding/json"
	"net/http"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"github.com/lapostoj/winemanager/service/presentation/api/request"
	"github.com/lapostoj/winemanager/service/presentation/api/response"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Test handles the GET calls to '/api/test'
func Test(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if err := persistence.TestWine(ctx); err != nil {
		log.Errorf(ctx, "TestWine - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("TestWine added"))
}

// GetWines handles the GET calls to '/api/wines' and return the stored wines (non 0 quantity)
func GetWines(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	w.Header().Set("Access-Control-Allow-Origin", Website)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var wines []wine.Wine
	if err := persistence.GetWinesInStock(ctx, &wines); err != nil {
		log.Errorf(ctx, "GetWines - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(response.NewWineResponses(wines))
	if err != nil {
		log.Errorf(ctx, "GetWines - marshal: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

// OptionsWines handles the OPTIONS calls to '/api/wines' and check their headers
func OptionsWines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", Website)
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// PostWines handles the POST calls to '/api/wines' and add the wine in the database
func PostWines(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	var postWineRequest request.PostWineRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postWineRequest)
	if err != nil {
		log.Errorf(ctx, "PostWines - decode: %q\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wine := postWineRequest.NewWine()
	key, err := persistence.SaveWine(ctx, wine)
	if err != nil {
		log.Errorf(ctx, "PostWines - persistence: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(key))
}
