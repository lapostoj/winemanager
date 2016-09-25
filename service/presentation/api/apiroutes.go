package api

import (
	"net/http"
)

const website = "http://cave-inventaire.appspot.com"

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
		Test,
	},
	Route{
		"GetWines",
		http.MethodGet,
		"/api/wines",
		GetWines,
	},
	Route{
		"OptionsWines",
		http.MethodOptions,
		"/api/wines",
		OptionsWines,
	},
	Route{
		"PostWines",
		http.MethodPost,
		"/api/wines",
		PostWines,
	},
	Route{
		"PostImport",
		http.MethodPost,
		"/api/import",
		PostImport,
	},
}
