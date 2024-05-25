package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/domain/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io/ioutil"
	"net/http"
)

type SendZipcode interface {
	Execute(ctx context.Context, zipcode *entity.ZipCode) (int, []byte, error)
}

type SendZipcodeUseCase struct {
	URL string
}

func (uc *SendZipcodeUseCase) Execute(ctx context.Context, zipcode *entity.ZipCode) (int, []byte, error) {
	requestURL := fmt.Sprintf("%s?zipcode=%s", uc.URL, zipcode.Code)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil, errors.New(string(body))
	}

	return resp.StatusCode, body, nil
}
