package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lapostoj/winemanager/service/application/service/createbottle"
	"github.com/lapostoj/winemanager/service/application/service/createcellar"
	"github.com/lapostoj/winemanager/service/application/service/createwine"
	"github.com/lapostoj/winemanager/service/application/service/getbottles"
	"github.com/lapostoj/winemanager/service/application/service/getcellar"
	"github.com/lapostoj/winemanager/service/application/service/getwines"
	persistence "github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"github.com/lapostoj/winemanager/service/presentation/api"
)

const defaultPort = "8080"

// main is called before the application starts.
func main() {
	frontendFolder := os.Getenv("FRONTEND_FOLDER")
	cellarRepository := persistence.CellarRepository{}
	getCellarService := getcellar.GetCellar{CellarRepository: cellarRepository}
	createCellarService := createcellar.CreateCellar{CellarRepository: cellarRepository}
	cellarHandler := api.CellarHandler{GetCellar: getCellarService, CreateCellar: createCellarService}

	wineRepository := persistence.WineRepository{}
	getWinesService := getwines.GetWines{WineRepository: wineRepository}
	createWineService := createwine.CreateWine{WineRepository: wineRepository}
	wineHandler := api.WineHandler{GetWines: getWinesService, CreateWine: createWineService}

	bottleRepository := persistence.BottleRepository{}
	getBottleService := getbottles.GetBottles{BottleRepository: bottleRepository}
	createBottleService := createbottle.CreateBottle{BottleRepository: bottleRepository}
	bottleHandler := api.BottleHandler{GetBottles: getBottleService, CreateBottle: createBottleService}

	router := api.NewRouter(cellarHandler, wineHandler, bottleHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(frontendFolder)))
	http.Handle("/api/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
