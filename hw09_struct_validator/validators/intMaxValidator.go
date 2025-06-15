package validators

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrIntMaxValidator = errors.New("value more than limitation")

type IntMaxValidator struct {
	AbstractValidator
}

func (s *IntMaxValidator) Validate() error {
	limitation, err := strconv.Atoi(s.limitation)
	actual := int(reflect.ValueOf(s.actual).Int())

	if err != nil {
		return s.parsingError("invalid format")
	}

	if actual > limitation {
		return ErrIntMaxValidator
	}

	return nil
}

func NewIntMaxValidator() Validator {
	return &IntMaxValidator{}
}
