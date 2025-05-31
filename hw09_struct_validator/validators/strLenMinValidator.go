package validators

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrStrLenMinValidator = errors.New("value less than limitation")

type StrLenMinValidator struct {
	AbstractValidator
}

func (s *StrLenMinValidator) Validate() error {
	limitation, err := strconv.Atoi(s.limitation)
	actual := len(reflect.ValueOf(s.actual).String())

	if err != nil {
		return s.parsingError("invalid format")
	}

	if limitation > 0 && actual < limitation {
		return ErrStrLenMinValidator
	}

	return nil
}

func NewStrLenMinValidator() Validator {
	return &StrLenMinValidator{}
}
