package fixer

import (
	"encoding/json"
	"fmt"
	"github.com/wintermute-core/currency-conversion/env"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Request - parameters to do request to Fixer API
type Request struct {
	ApiKey  string
	Timeout time.Duration
}

// Response - Fixer API response
type Response struct {
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}

func (fixer Request) FetchExchangeRate() (*Response, error) {
	if env.IsDefined("TRACE") {
		log.Println("Enter FetchExchangeRate")
		defer log.Println("Exit FetchExchangeRate")
	}
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s&format=1", fixer.ApiKey)
	spaceClient := http.Client{
		Timeout: fixer.Timeout,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return nil, readErr
	}

	data, err := parseResponse(body)
	return data, err

}

func parseResponse(body []byte) (*Response, error) {
	data := Response{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &data, nil
}
