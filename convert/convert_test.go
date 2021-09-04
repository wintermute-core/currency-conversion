package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyConversion(t *testing.T) {

	LoadRates("EUR", map[string]float64{
		"USD": 1.187905,
		"EUR": 1,
	})

	amount, err := Exchange(1, "EUR", "USD")
	assert.NoError(t, err)
	assert.Equal(t, 1.19, amount)

	amount, err = Exchange(2, "USD", "EUR")
	assert.NoError(t, err)
	assert.Equal(t, 1.68, amount)
}

func TestInvalidCurrencies(t *testing.T) {
	_, err := Exchange(1, "EUR", "BTC")
	assert.Error(t, err)

	for k := range exchangeRates {
		delete(exchangeRates, k)
	}

	_, err = Exchange(1, "EUR", "USD")
	assert.Error(t, err)
}
