package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lapostoj/winemanager/service/application/service"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"github.com/lapostoj/winemanager/service/presentation/api"
)

const defaultPort = "8080"

// main is called before the application starts.
func main() {
	frontend_folder := os.Getenv("FRONTEND_FOLDER")
	cellarRepository := persistence.CellarRepository{}
	getCellarService := service.GetCellar{Repository: cellarRepository}
	cellarHandler := api.CellarHandler{GetCellar: getCellarService}
	router := api.NewRouter(cellarHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(frontend_folder)))
	http.Handle("/api/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
