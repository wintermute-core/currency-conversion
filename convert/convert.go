package convert

import (
	"fmt"
	"github.com/wintermute-core/currency-conversion/env"
	"log"
	"strconv"
)

// allowedCurrencies - allowed currencies to exchange
var allowedCurrencies = map[string]bool{
	"EUR": true,
	"USD": true,
}

// baseCurrency - base currency for exchanges
var baseCurrency = "EUR"

// exchangeRates - available exchange rates
var exchangeRates = map[string]float64{}

func Exchange(amount float64, from string, to string) (float64, error) {
	if env.IsDefined("TRACE") {
		log.Printf("Enter exchange: %v, %v, %v \n", amount, from, to)
		defer log.Printf("Exit exchange: %v, %v, %v \n", amount, from, to)
	}

	if err := validateCurrency(from); err != nil {
		return 0, err
	}

	if err := validateCurrency(to); err != nil {
		return 0, err
	}

	if from == to {
		return 0, ExchangeCurrencyEqual{}
	}

	var value float64
	if from == baseCurrency {
		value = amount * exchangeRates[to]
	} else {
		value = amount * (1 / exchangeRates[from])
	}
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 10)
	if err != nil {
		return 0, err
	}
	return value, nil

}

// validateCurrency - check if passed currency is in allowed list and exchange rates are available
func validateCurrency(currency string) error {
	_, found := allowedCurrencies[currency]
	if !found {
		return InvalidExchangeCurrency{
			Currency: currency,
		}
	}
	_, found = exchangeRates[currency]
	if !found {
		return NoExchangeRates{
			Currency: currency,
		}
	}
	return nil
}

// LoadRates - load exchange rates
func LoadRates(newBaseCurrency string, newRates map[string]float64) {
	if env.IsDefined("TRACE") {
		log.Printf("Enter LoadRates: %v, %v\n", newBaseCurrency, newRates)
		defer log.Printf("Exit LoadRates: %v, %v\n", newBaseCurrency, newRates)
	}
	filteredRates := map[string]float64{}
	for key, _ := range allowedCurrencies {
		filteredRates[key] = newRates[key]
	}
	exchangeRates = filteredRates
	baseCurrency = newBaseCurrency
}

type ExchangeCurrencyEqual struct{}

func (m ExchangeCurrencyEqual) Error() string {
	return "Exchange currencies are equal"
}

type InvalidExchangeCurrency struct {
	Currency string
}

func (m InvalidExchangeCurrency) Error() string {
	return fmt.Sprintf("Invalid exchange currency: %s", m.Currency)
}

type NoExchangeRates struct {
	Currency string
}

func (m NoExchangeRates) Error() string {
	return fmt.Sprintf("No exchange rates for currency: %s", m.Currency)
}
