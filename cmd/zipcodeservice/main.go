package main

import (
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/interface/controller"
	usecase2 "github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/usecase"
	"log"
	"net/http"
)

func main() {
	var validateZipcodeUseCase usecase2.ValidateZipcode = &usecase2.ValidateZipcodeUseCase{}
	var sendZipcodeUseCase usecase2.SendZipcode = &usecase2.SendZipcodeUseCase{URL: "http://localhost:8081/weather"}

	zipcodeController := controller.NewZipcodeController(validateZipcodeUseCase, sendZipcodeUseCase)

	http.HandleFunc("/zipcode", zipcodeController.Handle)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
