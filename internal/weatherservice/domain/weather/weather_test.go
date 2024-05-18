package weather

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnInvalidCelsiusWhenNewLocationThenReturnError(t *testing.T) {
	// Given
	celsius := -300.0

	// When
	_, err := NewWeather(celsius)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, "invalid celsius", err.Error())
}

func TestGivenAValidParamsWhenNewLocationThenReturnLocationWithAllParams(t *testing.T) {
	// Given
	celsius := 10.0

	// When
	location, err := NewWeather(celsius)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, celsius, location.Celsius)
	assert.Equal(t, 50.0, location.Fahrenheit)
	assert.Equal(t, 283.15, location.Kelvin)
}

func TestGivenAValidCityAndCelsiusWhenConvertFahrenheitThenReturnFahrenheit(t *testing.T) {
	// Given
	celsius := 10.0
	location, _ := NewWeather(celsius)

	// When
	location.ConvertFahrenheit()

	// Then
	assert.Equal(t, 50.0, location.Fahrenheit)
}

func TestGivenAValidCityAndCelsiusWhenConvertKelvinThenReturnKelvin(t *testing.T) {
	// Given
	celsius := 10.0
	location, _ := NewWeather(celsius)

	// When
	location.ConvertKelvin()

	// Then
	assert.Equal(t, 283.15, location.Kelvin)
}
