package entity

import (
	"errors"
	"regexp"
)

type ZipCode struct {
	Code string
}

func NewZipCode(code string) (*ZipCode, error) {
	if err := validateZipCode(code); err != nil {
		return nil, err
	}
	return &ZipCode{Code: code}, nil
}

func validateZipCode(code string) error {
	match, _ := regexp.MatchString(`^\d{8}$`, code)
	if !match {
		return errors.New("invalid zip code")
	}
	return nil
}
