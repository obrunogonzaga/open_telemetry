package usecase

import (
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/domain/entity"
)

type ValidateZipcode interface {
	Execute(code string) (*entity.ZipCode, error)
}

type ValidateZipcodeUseCase struct{}

func (uc *ValidateZipcodeUseCase) Execute(code string) (*entity.ZipCode, error) {
	return entity.NewZipCode(code)
}
