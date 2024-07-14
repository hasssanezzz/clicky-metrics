package cmd

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the login route :O")
}
