package winemanager

import (
	"fmt"
	"net/http"

	"github.com/lapostoj/winemanager/service/presentation"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

// Init is called before the application starts.
func init() {
	router := presentation.NewRouter()
	http.HandleFunc("/", welcome)
	http.Handle("/api/", router)
}

// This will just ask the user to identify.
// In itself it doesn't add any restriction after this identification.
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")

	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}
