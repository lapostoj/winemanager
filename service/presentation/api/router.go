package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lapostoj/winemanager/service/application/service"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
)

// Route structure for API.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// GetClientURL returns the CLIENT_URL from the env variables or a default
func GetClientURL() string {
	if value, ok := os.LookupEnv("CLIENT_URL"); ok {
		return value
	}
	return "http://localhost:3000"
}

// NewRouter creates a router to mach routes and handlers
func NewRouter(cellarHandler CellarHanderInterface) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	routes := []Route{
		Route{
			"Welcome",
			http.MethodGet,
			"/api/test",
			WineHandler{WineRepository: persistence.WineRepository{}}.Test,
		},
		Route{
			"GetCellars",
			http.MethodGet,
			"/api/cellars",
			cellarHandler.QueryCellars,
		},
		Route{
			"PostCellar",
			http.MethodPost,
			"/api/cellars",
			cellarHandler.PostCellar,
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

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
