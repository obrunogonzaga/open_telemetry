package cep

import (
	"errors"
	"regexp"
)

func IsValid(cep string) error {
	cepRegex := regexp.MustCompile(`^\d{8}$`)
	if !cepRegex.MatchString(cep) {
		return errors.New("invalid zipcode")
	}
	return nil
}
