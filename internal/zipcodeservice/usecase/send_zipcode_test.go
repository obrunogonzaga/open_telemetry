package usecase

import (
	"github.com/obrunogonzaga/open-telemetry/internal/zipcodeservice/domain/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendZipcodeUseCase_Execute(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		uc := &SendZipcodeUseCase{URL: server.URL}
		zip := &entity.ZipCode{Code: "12345678"}
		status, _, err := uc.Execute(zip)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if status != http.StatusOK {
			t.Errorf("Expected status code to be %v, got %v", http.StatusOK, status)
		}
	})

	t.Run("unsuccessful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Internal Server Error"))
		}))
		defer server.Close()

		uc := &SendZipcodeUseCase{URL: server.URL}
		zip := &entity.ZipCode{Code: "12345678"}
		status, _, err := uc.Execute(zip)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if status != http.StatusInternalServerError {
			t.Errorf("Expected status code to be %v, got %v", http.StatusInternalServerError, status)
		}
		if err.Error() != "Internal Server Error" {
			t.Errorf("Expected error message to be 'Internal Server Error', got %v", err.Error())
		}
	})
}
