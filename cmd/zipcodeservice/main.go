package main

import (
	"github.com/obrunogonzaga/open-telemetry/configs"
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/interface/controller"
	usecase2 "github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/usecase"
	"log"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	var validateZipcodeUseCase usecase2.ValidateZipcode = &usecase2.ValidateZipcodeUseCase{}
	var sendZipcodeUseCase usecase2.SendZipcode = &usecase2.SendZipcodeUseCase{URL: config.WeatherServiceURL}

	zipcodeController := controller.NewZipcodeController(validateZipcodeUseCase, sendZipcodeUseCase)

	http.HandleFunc("/zipcode", zipcodeController.Handle)

	log.Println("Server running on port " + config.ZipCodeServerPort)
	log.Fatal(http.ListenAndServe(config.ZipCodeServerPort, nil))
}
