package usecase

import (
	"context"
	locationService "github.com/obrunogonzaga/open-telemetry/internal/weatherservice/service"
)

type Input struct {
	CEP string `json:"cep"`
}

type Output struct {
	CEP   string `json:"cep"`
	City  string `json:"localidade"`
	State string `json:"uf"`
}

type FindLocationUseCase struct {
	LocationService locationService.LocationService
}

func NewFindLocationUseCase(LocationService locationService.LocationService) *FindLocationUseCase {
	return &FindLocationUseCase{
		LocationService: LocationService,
	}
}

func (c *FindLocationUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	location, err := c.LocationService.FindLocationByZipCode(ctx, input.CEP)
	if err != nil {
		return Output{}, err
	}

	return Output{
		CEP:  location.CEP,
		City: location.City,
	}, nil
}
