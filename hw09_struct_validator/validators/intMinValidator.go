package validators

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrIntMinValidator = errors.New("value less than limitation")

type IntMinValidator struct {
	AbstractValidator
}

func (s *IntMinValidator) Validate() error {
	limitation, err := strconv.Atoi(s.limitation)
	actual := int(reflect.ValueOf(s.actual).Int())

	if err != nil {
		return s.parsingError("invalid format")
	}

	if actual < limitation {
		return ErrIntMinValidator
	}

	return nil
}

func NewIntMinValidator() Validator {
	return &IntMinValidator{}
}
