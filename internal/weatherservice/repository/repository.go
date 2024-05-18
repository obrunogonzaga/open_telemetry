package repository

import (
	"context"
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/location"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/weather"
)

type WeatherRepository interface {
	GetWeather(ctx context.Context, city string, config configs.Config) (*weather.Weather, error)
}

type LocationRepository interface {
	FindCityByZipCode(ctx context.Context, cep string) (*location.Location, error)
}
