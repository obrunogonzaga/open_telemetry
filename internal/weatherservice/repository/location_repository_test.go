package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/location"
	"github.com/stretchr/testify/assert"
)

func TestFindCityByZipCode(t *testing.T) {
	// Create a test server that returns a predefined response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep": "01153000", "city": "São Paulo"}`))
	}))
	defer ts.Close()

	// Create a new LocationRepository with a mocked HTTP client
	repo := NewLocationRepository(ts.Client())

	// Call the function with a test zip code
	loc, err := repo.FindCityByZipCode(context.Background(), "01153000")

	// Assert that there was no error and the response was correctly parsed
	assert.NoError(t, err)
	assert.Equal(t, &location.Location{
		CEP:  "01153-000",
		City: "São Paulo",
	}, loc)
}

func TestFindCityByInvalifZipCode(t *testing.T) {
	// Create a test server that returns a predefined response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep": "0115300A", "city": "São Paulo"}`))
	}))
	defer ts.Close()

	// Create a new LocationRepository with a mocked HTTP client
	repo := NewLocationRepository(ts.Client())

	// Call the function with a test zip code
	_, err := repo.FindCityByZipCode(context.Background(), "0115300A")

	// Assert that there was no error and the response was correctly parsed
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid zipcode")
}
