package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/lapostoj/winemanager/service/application/service"
	"github.com/lapostoj/winemanager/service/infrastructure/utils"

	"github.com/lapostoj/winemanager/service/presentation/api/response"
)

type CellarHanderInterface interface {
	QueryCellars(w http.ResponseWriter, r *http.Request)
}

// CellarHandler implements handling of api calls for cellars
type CellarHandler struct {
	GetCellar service.GetCellarService
}

// QueryCellars handles the GET calls to '/api/cellars' and return the cellars matching the query
func (handler CellarHandler) QueryCellars(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", Website)
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
