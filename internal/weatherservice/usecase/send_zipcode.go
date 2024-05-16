package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/entity"
	"io/ioutil"
	"net/http"
)

type SendZipcode interface {
	Execute(zipcode *entity.ZipCode) (int, error)
}

type SendZipcodeUseCase struct {
	URL string
}

func (uc *SendZipcodeUseCase) Execute(zipcode *entity.ZipCode) (int, error) {
	data := map[string]string{"cep": zipcode.Code}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp, err := http.Post(uc.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return resp.StatusCode, errors.New(string(body))
	}

	return resp.StatusCode, nil
}
