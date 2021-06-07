package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()
	router.Handle("/ping", http.HandlerFunc(httPing))
	return router
}

func httPing(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "pong")
}
