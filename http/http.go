package http

import (
	"github.com/denis256/currency-conversion/env"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Server() error {
	listenPort := env.Env("HTTP_PORT", "8080")
	log.Printf("Starting HTTP server on %v", listenPort)
	r := router()
	err := http.ListenAndServe(":"+listenPort, r)
	if err != nil {
		return err
	}
	return nil
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	return router
}
