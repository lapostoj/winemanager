package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/lapostoj/winemanager/service/presentation/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCellarHandler struct {
	mock.Mock
}

func (mock MockCellarHandler) QueryCellars(w http.ResponseWriter, r *http.Request) {
}

func (mock MockCellarHandler) PostCellar(w http.ResponseWriter, r *http.Request) {
}

type MockWineHandler struct {
	mock.Mock
}

func (mock MockWineHandler) QueryWines(w http.ResponseWriter, r *http.Request) {
}

func (mock MockWineHandler) OptionsWines(w http.ResponseWriter, r *http.Request) {
}

func (mock MockWineHandler) PostWine(w http.ResponseWriter, r *http.Request) {
}

func (mock MockWineHandler) PostTest(w http.ResponseWriter, r *http.Request) {
}

func TestNewRouterHandlesPostImport(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("POST", "/api/import", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "PostImport", match.Route.GetName())
}

func TestNewRouterHandlesGetCellar(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("GET", "/api/cellars", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "GetCellars", match.Route.GetName())
}

func TestNewRouterHandlesPostCellar(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("POST", "/api/cellars", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "PostCellar", match.Route.GetName())
}

func TestNewRouterHandlesGetWines(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("GET", "/api/wines", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "GetWines", match.Route.GetName())
}

func TestNewRouterHandlesOptionsWines(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("OPTIONS", "/api/wines", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "OptionsWines", match.Route.GetName())
}

func TestNewRouterHandlesPostWine(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("POST", "/api/wines", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "PostWine", match.Route.GetName())
}

func TestNewRouterHandlesPostTest(t *testing.T) {
	var body bytes.Buffer
	var match mux.RouteMatch
	cellarHandler := new(MockCellarHandler)
	wineHandler := new(MockWineHandler)
	router := api.NewRouter(cellarHandler, wineHandler)
	request := httptest.NewRequest("POST", "/api/test", &body)

	assert.True(t, router.Match(request, &match))
	assert.Equal(t, "PostTest", match.Route.GetName())
}
