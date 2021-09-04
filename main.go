package main

import (
	"github.com/denis256/currency-conversion/http"
	"github.com/denis256/currency-conversion/sync"
)

func main() {

	err := sync.FetchCurrencies()
	if err != nil {
		panic(err)
	}

	err = http.Server()
	if err != nil {
		panic(err)
	}
}
