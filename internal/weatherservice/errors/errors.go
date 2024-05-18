package errors

import "errors"

var ErrInvalidCEP = errors.New("invalid zipcode")
var ErrZipCodetNotFound = errors.New("can not find zipcode")
