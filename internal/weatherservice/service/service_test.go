package service

import (
	"context"
	"errors"
	"testing"

	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/location"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLocationRepository struct {
	mock.Mock
}

func (m *MockLocationRepository) FindCityByZipCode(ctx context.Context, cep string) (*location.Location, error) {
	args := m.Called(ctx, cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*location.Location), args.Error(1)
}

func TestFindLocationByZipCode(t *testing.T) {
	mockRepo := new(MockLocationRepository)
	mockRepo.On("FindCityByZipCode", mock.Anything, "01001-000").Return(&location.Location{CEP: "01001-000", City: "São Paulo"}, nil)
	mockRepo.On("FindCityByZipCode", mock.Anything, "invalid").Return(nil, errors.New("invalid zip code"))

	locationService := NewLocationService(mockRepo)

	locationByZipCode, err := locationService.FindLocationByZipCode(context.Background(), "01001-000")
	assert.NoError(t, err)
	assert.Equal(t, &location.Location{CEP: "01001-000", City: "São Paulo"}, locationByZipCode)

	_, err = locationService.FindLocationByZipCode(context.Background(), "invalid")
	assert.Error(t, err)
}
