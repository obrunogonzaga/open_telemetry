package repository

import (
	"context"
	"encoding/json"
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock output for successful weather API response
var mockOutput = Output{
	Location: struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	}{
		Name: "London",
	},
	Current: struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	}{
		TempC: 20.0,
	},
}

// MockWeatherAPI creates a mock server for testing
func MockWeatherAPI(t *testing.T) (*httptest.Server, *WeatherAPI) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(mockOutput); err != nil {
			t.Fatalf("Failed to encode mock output: %v", err)
		}
	})

	server := httptest.NewServer(handler)
	client := server.Client()

	return server, NewWeatherAPI(client)
}

func TestGetWeather(t *testing.T) {
	server, api := MockWeatherAPI(t)
	defer server.Close()

	ctx := context.Background()
	config := configs.Config{WeatherApiKey: "dc615a7639ce41dc862232340242504"}

	weather, err := api.GetWeather(ctx, "London", config)
	assert.NoError(t, err)
	assert.NotNil(t, weather)
}
