package controller

import (
	"encoding/json"
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/usecase"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type ZipcodeController struct {
	validateZipcode usecase.ValidateZipcode
	sendZipcode     usecase.SendZipcode
	tracer          trace.Tracer
}

func NewZipcodeController(v usecase.ValidateZipcode, s usecase.SendZipcode, t trace.Tracer) *ZipcodeController {
	return &ZipcodeController{
		validateZipcode: v,
		sendZipcode:     s,
		tracer:          t,
	}
}

func (c *ZipcodeController) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, span := c.tracer.Start(r.Context(), "HandleZipcodeRequest")
	defer span.End()

	var req struct {
		Zipcode string `json:"zipcode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	zipcode, err := c.validateZipcode.Execute(req.Zipcode)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	status, body, err := c.sendZipcode.Execute(ctx, zipcode)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
