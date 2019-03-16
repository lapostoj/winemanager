package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/lapostoj/winemanager/service/presentation"
)

const defaultPort = "8080"

// Init is called before the application starts.
func main() {
	router := presentation.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))
	http.Handle("/api/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

// This will just ask the user to identify.
// In itself it doesn't add any restriction after this identification.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// u := user.Current(ctx)
	// if u == nil {
	// 	url, _ := user.LoginURL(ctx, "/")
	// 	fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
	// 	return
	// }
	//url, _ := user.LogoutURL(ctx, "/")
	//fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)

	filePath := "./app/index.html"
	if r.URL.Path == "/" {
		filePath = "./app/index.html"
	} else {
		filePath = "./app" + r.URL.Path
	}
	file, err := os.Open(filePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fileByte, notFoundErr := ioutil.ReadFile("./app/notfound.html")
		if notFoundErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `<html><body style="font-size: 50px">Erreur...</body></html>`)
		}
		fmt.Fprint(w, string(fileByte))
		return
	}

	defer file.Close()
	_, filename := path.Split(filePath)
	http.ServeContent(w, r, filename, time.Time{}, file)
}
