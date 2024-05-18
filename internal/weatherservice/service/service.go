package service

import (
	"context"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/location"
)

type LocationService interface {
	FindLocationByZipCode(ctx context.Context, cep string) (*location.Location, error)
}
