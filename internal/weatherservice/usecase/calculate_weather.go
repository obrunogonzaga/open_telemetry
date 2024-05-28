package usecase

import (
	"context"
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/repository"
)

type CalculateWeatherInput struct {
	City 		string `json:"city"`
}

type CalculateWeatherOutput struct {
	City	   string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type CalculateWeatherUseCase struct {
	Gateway repository.WeatherRepository
	Config  *configs.Config
}

func NewCalculateWeatherUseCase(Gateway repository.WeatherRepository, Config *configs.Config) *CalculateWeatherUseCase {
	return &CalculateWeatherUseCase{
		Gateway: Gateway,
		Config:  Config,
	}
}

func (c *CalculateWeatherUseCase) Execute(ctx context.Context, input CalculateWeatherInput) (CalculateWeatherOutput, error) {
	weather, err := c.Gateway.GetWeather(ctx, input.City, *c.Config)
	if err != nil {
		return CalculateWeatherOutput{}, err
	}

	return CalculateWeatherOutput{
		City:       input.City,
		Celsius:    weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin:     weather.Kelvin,
	}, nil
}
