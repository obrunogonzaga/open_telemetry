package errors

import (
	"testing"
)

func TestErrors(t *testing.T) {
	t.Run("ErrInvalidCEP", func(t *testing.T) {
		err := ErrInvalidCEP
		if err == nil || err.Error() != "invalid zipcode" {
			t.Errorf("ErrInvalidCEP error message expected to be 'invalid zipcode', got %v", err)
		}
	})

	t.Run("ErrZipCodetNotFound", func(t *testing.T) {
		err := ErrZipCodetNotFound
		if err == nil || err.Error() != "can not find zipcode" {
			t.Errorf("ErrZipCodetNotFound error message expected to be 'can not find zipcode', got %v", err)
		}
	})
}
