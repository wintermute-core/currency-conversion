package http

import (
	"encoding/json"
	"github.com/wintermute-core/currency-conversion/convert"
	"github.com/wintermute-core/currency-conversion/env"
	"github.com/wintermute-core/currency-conversion/project"
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
	router.HandleFunc("/api/v1/public/project/new", NewProject).Methods("POST")

	securedPath := router.PathPrefix("/api/v1/private").Subrouter()

	securedPath.Use(ApiKeyVerify)
	securedPath.HandleFunc("/exchange", Exchange).Methods("POST")

	return router
}

func NewProject(w http.ResponseWriter, r *http.Request) {
	if env.IsDefined("TRACE") {
		log.Println("Enter NewProject")
		defer log.Println("Exit NewProject")
	}

	newProject := project.NewProject()
	err := json.NewEncoder(w).Encode(newProject)
	if err != nil {
		handleError(w, err)
	}
	return
}

func Exchange(w http.ResponseWriter, r *http.Request) {
	if env.IsDefined("TRACE") {
		log.Println("Enter Exchange")
		defer log.Println("Exit Exchange")
	}
	var exchangeRequest = ExchangeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&exchangeRequest); err != nil {
		handleError(w, err)
		return
	}
	amount, err := convert.Exchange(exchangeRequest.Amount, exchangeRequest.From, exchangeRequest.To)
	if err != nil {
		handleError(w, err)
		return
	}
	var exchangeResponse = ExchangeResponse{Amount: amount}
	if err := json.NewEncoder(w).Encode(exchangeResponse); err != nil {
		handleError(w, err)
		return
	}
	return
}

func ApiKeyVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if env.IsDefined("TRACE") {
			log.Println("Enter ApiKeyVerify")
			defer log.Println("Exit ApiKeyVerify")
		}

		var header = r.Header.Get("x-api-key")
		if !project.IsValidApiKey(header) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("Error in creating new project %v", err)
	if _, errw := w.Write([]byte(err.Error())); errw != nil {
		log.Println(errw)
	}
}

// ExchangeRequest - request for exchange
type ExchangeRequest struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

// ExchangeResponse - exchange response
type ExchangeResponse struct {
	Amount float64 `json:"amount"`
}
