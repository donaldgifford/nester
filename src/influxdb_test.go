package src_test

import (
	"testing"

	"github.com/donaldgifford/nester/src"
	"github.com/stretchr/testify/assert"
)

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	c := src.ConvertCelsiusToFahrenheit(100)
	assert.Equal(t, c, 212.0)
}

func TestConvertStatusToBoolOnline(t *testing.T) {
	s := src.ConvertStatusToBool("ONLINE")
	assert.Equal(t, s, true)
}

func TestConvertStatusToBoolOffline(t *testing.T) {
	s := src.ConvertStatusToBool("OFFLINE")
	assert.Equal(t, s, false)
}
