package main

import (
	"context"
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/infra/web"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/infra/web/webserver"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/repository"
	locatiionService "github.com/obrunogonzaga/open-telemetry/internal/weatherservice/service"
	"github.com/obrunogonzaga/open-telemetry/pkg/tracing"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Graceful shutdown - begin
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	// Graceful shutdown - end

	shutdown, err := tracing.InitProvider(ctx, "weather-service", os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TraceProvider", err)
		}
	}()

	tracer := otel.Tracer("weather-service")

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	locationRepo := repository.NewLocationRepository(client)
	locationService := locatiionService.NewLocationService(locationRepo)
	weather := repository.NewWeatherAPI(client)
	//TODO: Passar o tracer para o handler
	handler := web.NewHandler(locationService, weather, config, tracer)

	//TODO: Implementar a injeção de dependência com o wire
	restServer := webserver.NewWebServer(config.WebServerPort)

	restServer.AddHandler("/weather", handler.Execute)

	go func() {
		if err := restServer.Start(); err != nil {
			log.Fatal("failed to start server", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}

	// Create a timeout context for the graceful shutdown
	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

}
