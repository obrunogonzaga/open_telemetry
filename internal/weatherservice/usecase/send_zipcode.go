package usecase

import (
	"errors"
	"fmt"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/entity"
	"io/ioutil"
	"net/http"
)

type SendZipcode interface {
	Execute(zipcode *entity.ZipCode) (int, []byte, error)
}

type SendZipcodeUseCase struct {
	URL string
}

func (uc *SendZipcodeUseCase) Execute(zipcode *entity.ZipCode) (int, []byte, error) {
	requestURL := fmt.Sprintf("%s?zipcode=%s", uc.URL, zipcode.Code)

	resp, err := http.Get(requestURL)
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
