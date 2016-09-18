package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/infrastructure/persistence/appengine"
	"github.com/lapostoj/winemanager/service/presentation/api/response"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const website = "http://cave-inventaire.appspot.com/"

// Test handles the GET calls to '/api/test'
func Test(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if err := persistence.TestWine(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "TestWine added")
}

// GetWines handles the GET calls to '/api/wines' and return the stored wines (non 0 quantity)
func GetWines(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var wines []wine.Wine
	if err := persistence.GetWinesInStock(ctx, &wines); err != nil {
		log.Errorf(ctx, "GetWines: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := json.Marshal(response.NewWineResponses(wines))
	if err != nil {
		log.Errorf(ctx, "GetWines: %q\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(response))
}
