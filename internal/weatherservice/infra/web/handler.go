package web

import (
	"encoding/json"
	"errors"
	"github.com/obrunogonzaga/open-telemetry/configs"
	customErrors "github.com/obrunogonzaga/open-telemetry/internal/weatherservice/errors"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/repository"
	locationService "github.com/obrunogonzaga/open-telemetry/internal/weatherservice/service"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type Handler struct {
	LocationService locationService.LocationService
	WeatherService  repository.WeatherRepository
	Config          *configs.Config
	OTELTracer      trace.Tracer
}

func NewHandler(LocationService locationService.LocationService, WeatherService repository.WeatherRepository, Config *configs.Config, trace trace.Tracer) *Handler {
	return &Handler{
		LocationService: LocationService,
		WeatherService:  WeatherService,
		Config:          Config,
		OTELTracer:      trace,
	}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := h.OTELTracer.Start(ctx, "weather-service")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")

	zipCodeDTO := usecase.Input{
		CEP: r.URL.Query().Get("zipcode"),
	}

	findLocation := usecase.NewFindLocationUseCase(h.LocationService)
	output, err := findLocation.Execute(r.Context(), zipCodeDTO)
	if err != nil {
		if errors.Is(err, customErrors.ErrInvalidCEP) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, customErrors.ErrZipCodetNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherUseCase := usecase.NewCalculateWeatherUseCase(h.WeatherService, h.Config)
	weatherOutput, err := weatherUseCase.Execute(r.Context(), usecase.CalculateWeatherInput{City: output.City})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(weatherOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
