package fixer

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"

	"testing"
)

func TestParseFixerResponse(t *testing.T) {

	buf, err := ioutil.ReadFile("../../docs/fixer-io-example.json")
	if err != nil {
		panic(err)
	}
	response, err := parseResponse(buf)
	if err != nil {
		panic(err)
	}
	assert.True(t, response.Success)
	assert.Equal(t, "EUR", response.Base)
	assert.Equal(t, 1.18835, response.Rates["USD"])
}

func TestParseFailureResponse(t *testing.T) {

	buf, err := parseResponse([]byte("random text"))
	assert.Nil(t, buf)
	assert.NotNil(t, err)
}
