package service

import (
	"context"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/location"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/repository"
)

type locationServiceImpl struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationServiceImpl{
		repo: repo,
	}
}

func (v *locationServiceImpl) FindLocationByZipCode(ctx context.Context, cep string) (*location.Location, error) {
	return v.repo.FindCityByZipCode(ctx, cep)
}
