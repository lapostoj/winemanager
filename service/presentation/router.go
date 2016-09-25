package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lapostoj/winemanager/service/presentation/api"
)

// NewRouter creates a router to mach routes and handlers
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range api.APIRoutes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
