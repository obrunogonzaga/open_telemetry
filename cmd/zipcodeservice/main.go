package main

import (
	"context"
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/interface/controller"
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/usecase"
	"github.com/obrunogonzaga/open-telemetry/pkg/tracing"
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

	shutdown, err := tracing.InitProvider(ctx, "zipcode-service", os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TraceProvider", err)
		}
	}()

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	var validateZipcodeUseCase usecase.ValidateZipcode = &usecase.ValidateZipcodeUseCase{}
	var sendZipcodeUseCase usecase.SendZipcode = &usecase.SendZipcodeUseCase{URL: config.WeatherServiceURL}
	zipcodeController := controller.NewZipcodeController(validateZipcodeUseCase, sendZipcodeUseCase)

	http.HandleFunc("/zipcode", zipcodeController.Handle)
	log.Printf("Server running on port %s", config.ZipCodeServerPort)
	srv := &http.Server{
		Addr:    ":" + config.ZipCodeServerPort,
		Handler: nil,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
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
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")

}
