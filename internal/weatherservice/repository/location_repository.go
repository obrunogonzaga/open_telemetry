package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/location"
	customErrors "github.com/obrunogonzaga/open-telemetry/internal/weatherservice/errors"
	validadeZipCode "github.com/obrunogonzaga/open-telemetry/pkg/cep"
	"net/http"
)

type LocationOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type LocationRepositoryImpl struct {
	client *http.Client
}

func NewLocationRepository(client *http.Client) LocationRepository {
	return &LocationRepositoryImpl{
		client: client,
	}
}

func (v *LocationRepositoryImpl) FindCityByZipCode(ctx context.Context, cep string) (*location.Location, error) {
	err := validadeZipCode.IsValid(cep)
	if err != nil {
		return nil, customErrors.ErrInvalidCEP
	}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var output LocationOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}

	if output.Cep == "" {
		return nil, customErrors.ErrZipCodetNotFound
	}

	return &location.Location{
		CEP:  output.Cep,
		City: output.Localidade,
	}, nil
}
