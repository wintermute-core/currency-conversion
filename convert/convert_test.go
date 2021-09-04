package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyConversion(t *testing.T) {

	exchangeRates["USD"] = 1.187905
	exchangeRates["EUR"] = 1

	amount, err := exchange(1, "EUR", "USD")
	assert.NoError(t, err)
	assert.Equal(t, 1.1879, amount)

	amount, err = exchange(2, "USD", "EUR")
	assert.NoError(t, err)
	assert.Equal(t, 1.6836, amount)
}

func TestInvalidCurrencies(t *testing.T) {
	_, err := exchange(1, "EUR", "BTC")
	assert.Error(t, err)

	for k := range exchangeRates {
		delete(exchangeRates, k)
	}

	_, err = exchange(1, "EUR", "USD")
	assert.Error(t, err)
}
