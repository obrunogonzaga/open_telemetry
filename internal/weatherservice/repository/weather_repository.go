package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/weather"
	"net/http"
	"net/url"
)

type Output struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
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
	} `json:"current"`
}

type WeatherAPI struct {
	Client *http.Client
}

func NewWeatherAPI(client *http.Client) *WeatherAPI {
	return &WeatherAPI{
		Client: client,
	}
}

func (w *WeatherAPI) GetWeather(ctx context.Context, city string, config configs.Config) (*weather.Weather, error) {
	//url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=" + config.WeatherApiKey + "&q=" + city + "&aqi=no")
	escapedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=" + "dc615a7639ce41dc862232340242504" + "&q=" + escapedCity + "&aqi=no")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var output Output
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}

	temperature, err := weather.NewWeather(output.Current.TempC)
	if err != nil {
		return nil, err
	}

	return temperature, nil
}
