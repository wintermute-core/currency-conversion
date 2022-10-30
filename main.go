package main

import (
	"log"
	"strconv"

	"github.com/jasonlvhit/gocron"
	"github.com/wintermute-core/currency-conversion/env"
	"github.com/wintermute-core/currency-conversion/http"
	"github.com/wintermute-core/currency-conversion/sync"
)

func main() {

	err := sync.FetchCurrencies()
	if err != nil {
		panic(err)
	}
	go executeSyncJob()

	err = http.Server()
	if err != nil {
		panic(err)
	}
}

func syncCurrenciesInBackground() {
	err := sync.FetchCurrencies()
	if err != nil {
		log.Println(err)
	}
}

func executeSyncJob() {
	syncInterval, err := strconv.Atoi(env.Env("SYNC_INTERVAL_MIN", "60"))
	if err != nil {
		panic(err)
	}
	err = gocron.Every(uint64(syncInterval)).Minutes().Do(syncCurrenciesInBackground)
	if err != nil {
		panic(err)
	}
	<-gocron.Start()
}
