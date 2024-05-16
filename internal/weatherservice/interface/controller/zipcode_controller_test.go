package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/domain/entity"
	"github.com/obrunogonzaga/open-telemetry/internal/weatherservice/interface/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for ValidateZipcode
type MockValidateZipcode struct {
	mock.Mock
}

func (m *MockValidateZipcode) Execute(code string) (*entity.ZipCode, error) {
	args := m.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.ZipCode), args.Error(1)
}

// Mock for SendZipcode
type MockSendZipcode struct {
	mock.Mock
}

func (m *MockSendZipcode) Execute(zipcode *entity.ZipCode) (int, error) {
	args := m.Called(zipcode)
	return args.Int(0), args.Error(1)
}

func TestHandle_ValidRequest(t *testing.T) {
	mockValidateZipcode := new(MockValidateZipcode)
	mockSendZipcode := new(MockSendZipcode)

	zipcode := &entity.ZipCode{Code: "12345"}
	mockValidateZipcode.On("Execute", "12345").Return(zipcode, nil)
	mockSendZipcode.On("Execute", zipcode).Return(http.StatusOK, nil)

	controller := controller.NewZipcodeController(mockValidateZipcode, mockSendZipcode)

	reqBody := map[string]string{"cep": "12345"}
	jsonReqBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/zipcode", bytes.NewBuffer(jsonReqBody))
	rr := httptest.NewRecorder()

	controller.Handle(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockValidateZipcode.AssertExpectations(t)
	mockSendZipcode.AssertExpectations(t)
}

func TestHandle_InvalidRequest(t *testing.T) {
	mockValidateZipcode := new(MockValidateZipcode)
	mockSendZipcode := new(MockSendZipcode)

	controller := controller.NewZipcodeController(mockValidateZipcode, mockSendZipcode)

	req, _ := http.NewRequest(http.MethodPost, "/zipcode", bytes.NewBuffer([]byte("{invalid-json")))
	rr := httptest.NewRecorder()

	controller.Handle(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandle_InvalidZipcode(t *testing.T) {
	mockValidateZipcode := new(MockValidateZipcode)
	mockSendZipcode := new(MockSendZipcode)

	mockValidateZipcode.On("Execute", "invalid").Return(nil, errors.New("invalid zipcode"))

	controller := controller.NewZipcodeController(mockValidateZipcode, mockSendZipcode)

	reqBody := map[string]string{"cep": "invalid"}
	jsonReqBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/zipcode", bytes.NewBuffer(jsonReqBody))
	rr := httptest.NewRecorder()

	controller.Handle(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	mockValidateZipcode.AssertExpectations(t)
}

func TestHandle_SendZipcodeError(t *testing.T) {
	mockValidateZipcode := new(MockValidateZipcode)
	mockSendZipcode := new(MockSendZipcode)

	zipcode := &entity.ZipCode{Code: "12345"}
	mockValidateZipcode.On("Execute", "12345").Return(zipcode, nil)
	mockSendZipcode.On("Execute", zipcode).Return(http.StatusInternalServerError, errors.New("internal error"))

	controller := controller.NewZipcodeController(mockValidateZipcode, mockSendZipcode)

	reqBody := map[string]string{"cep": "12345"}
	jsonReqBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/zipcode", bytes.NewBuffer(jsonReqBody))
	rr := httptest.NewRecorder()

	controller.Handle(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockValidateZipcode.AssertExpectations(t)
	mockSendZipcode.AssertExpectations(t)
}
