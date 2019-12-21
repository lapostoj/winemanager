package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
func NewRouter(
	cellarHandler CellarHanderInterface,
	wineHandler WineHandlerInterface,
	bottleHandler BottleHandlerInterface,
	importHandler ImportHandlerInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	routes := []Route{
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
			wineHandler.QueryWines,
		},
		Route{
			"OptionsWines",
			http.MethodOptions,
			"/api/wines",
			wineHandler.OptionsWines,
		},
		Route{
			"PostWine",
			http.MethodPost,
			"/api/wines",
			wineHandler.PostWine,
		},
		Route{
			"PostTest",
			http.MethodPost,
			"/api/test",
			wineHandler.PostTest,
		},
		Route{
			"GetBottles",
			http.MethodGet,
			"/api/bottles",
			bottleHandler.QueryBottles,
		},
		Route{
			"PostBottle",
			http.MethodPost,
			"/api/bottles",
			bottleHandler.PostBottle,
		},
		Route{
			"PostImport",
			http.MethodPost,
			"/api/import",
			importHandler.PostImport,
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
