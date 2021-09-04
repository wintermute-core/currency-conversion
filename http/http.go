package http

import (
	"github.com/denis256/currency-conversion/env"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Server() {
	listenPort := env.Env("HTTP_PORT", "8080")
	log.Printf("Starting HTTP server on %v", listenPort)
	r := router()
	http.ListenAndServe(":"+listenPort, r)

}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	return router
}
