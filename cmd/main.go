package winemanager

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"

	"github.com/lapostoj/winemanager/service/presentation"
)

// Init is called before the application starts.
func init() {
	router := presentation.NewRouter()
	http.HandleFunc("/", mainHandler)
	http.Handle("/api/", router)
}

// This will just ask the user to identify.
// In itself it doesn't add any restriction after this identification.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	//url, _ := user.LogoutURL(ctx, "/")

	http.ServeFile(w, r, "app"+r.URL.Path)
	//fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}
