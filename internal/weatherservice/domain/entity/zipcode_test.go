package entity

import (
	"testing"
)

func TestNewZipCode(t *testing.T) {
	t.Run("valid zip code", func(t *testing.T) {
		zip, err := NewZipCode("12345678")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if zip.Code != "12345678" {
			t.Errorf("Expected zip code to be '12345678', got %v", zip.Code)
		}
	})

	t.Run("invalid zip code", func(t *testing.T) {
		_, err := NewZipCode("1234")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("invalid zip code - invalid character", func(t *testing.T) {
		_, err := NewZipCode("1234567a")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
