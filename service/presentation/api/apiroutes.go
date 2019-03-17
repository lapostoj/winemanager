package api

import (
	"net/http"

	"github.com/lapostoj/winemanager/service/application/service"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
)

// Website is the url expected to use for the Access-Control-Allow-Origin header
const Website = "https://cave-inventaire.appspot.com"

// Route structure for API.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the specific routes of this API.
type Routes []Route

// APIRoutes defines the API routes for the winemanager
var APIRoutes = Routes{
	Route{
		"Welcome",
		http.MethodGet,
		"/api/test",
		WineHandler{WineRepository: persistence.WineRepository{}}.Test,
	},
	Route{
		"GetWines",
		http.MethodGet,
		"/api/wines",
		WineHandler{WineRepository: persistence.WineRepository{}}.GetWines,
	},
	Route{
		"OptionsWines",
		http.MethodOptions,
		"/api/wines",
		WineHandler{WineRepository: persistence.WineRepository{}}.OptionsWines,
	},
	Route{
		"PostWines",
		http.MethodPost,
		"/api/wines",
		WineHandler{WineRepository: persistence.WineRepository{}}.PostWines,
	},
	Route{
		"PostImport",
		http.MethodPost,
		"/api/import",
		ImportHandler{CsvImport: service.CsvImport{WineRepository: persistence.WineRepository{}}}.PostImport,
	},
}
