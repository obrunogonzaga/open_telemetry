package controller

import (
	"encoding/json"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/usecase"
	"net/http"
)

type ZipcodeController struct {
	validateZipcode usecase.ValidateZipcode
	sendZipcode     usecase.SendZipcode
}

func NewZipcodeController(v usecase.ValidateZipcode, s usecase.SendZipcode) *ZipcodeController {
	return &ZipcodeController{
		validateZipcode: v,
		sendZipcode:     s,
	}
}

func (c *ZipcodeController) Handle(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CEP string `json:"cep"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	zipcode, err := c.validateZipcode.Execute(req.CEP)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	status, err := c.sendZipcode.Execute(zipcode)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
}
