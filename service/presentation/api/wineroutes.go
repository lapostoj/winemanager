package api

import (
	"net/http"
)

// Route structure for API.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the specific routes of this API.
type Routes []Route

// WineRoutes defines the API routes for wine
var WineRoutes = Routes{
	Route{
		"Welcome",
		http.MethodGet,
		"/api/test",
		Test,
	},
}
