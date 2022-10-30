package sync

import (
	"log"
	"time"

	"github.com/wintermute-core/currency-conversion/api/fixer"
	"github.com/wintermute-core/currency-conversion/convert"
	"github.com/wintermute-core/currency-conversion/env"
)

const FixerApiKeyName = "FIXER_API_KEY"

// FetchCurrencies - query exchange rate service and save results in application
func FetchCurrencies() error {
	if env.IsDefined("TRACE") {
		log.Printf("Enter FetchCurrencies\n")
		defer log.Printf("Exit FetchCurrencies\n")
	}
	if !env.IsDefined(FixerApiKeyName) {
		return MissingFixerApiKey{}
	}
	apiKey := env.Env(FixerApiKeyName, "")
	request := fixer.Request{
		ApiKey:  apiKey,
		Timeout: time.Second * 1,
	}
	response, err := request.FetchExchangeRate()
	if err != nil {
		return err
	}
	convert.LoadRates(response.Base, response.Rates)

	return nil

}

type MissingFixerApiKey struct{}

func (m MissingFixerApiKey) Error() string {
	return "Missing Fixer API key: " + FixerApiKeyName
}
