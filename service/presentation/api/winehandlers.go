package api

import (
	"fmt"
	"net/http"
)

const website = "http://cave-inventaire.appspot.com/"

// Test handle the calls to '/api/test'
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to cave inventaire!")
}
