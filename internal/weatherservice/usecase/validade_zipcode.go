package usecase

import "github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/entity"

type ValidateZipcodeUseCase struct{}

func (uc *ValidateZipcodeUseCase) Execute(code string) (*entity.ZipCode, error) {
	return entity.NewZipCode(code)
}
