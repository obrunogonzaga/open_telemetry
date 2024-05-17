package main

import (
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/interface/controller"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/usecase"
	"log"
	"net/http"
)

func main() {
	var validateZipcodeUseCase usecase.ValidateZipcode = &usecase.ValidateZipcodeUseCase{}
	var sendZipcodeUseCase usecase.SendZipcode = &usecase.SendZipcodeUseCase{URL: "https://cloudrun-goexpert-za5o6n5xla-uc.a.run.app/weather"}

	zipcodeController := controller.NewZipcodeController(validateZipcodeUseCase, sendZipcodeUseCase)

	http.HandleFunc("/zipcode", zipcodeController.Handle)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
