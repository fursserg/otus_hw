package validators

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrStrLenMaxValidator = errors.New("value more than limitation")

type StrLenMaxValidator struct {
	AbstractValidator
}

func (s *StrLenMaxValidator) Validate() error {
	limitation, err := strconv.Atoi(s.limitation)
	actual := len(reflect.ValueOf(s.actual).String())

	if err != nil {
		return s.parsingError("invalid format")
	}

	if limitation > 0 && actual > limitation {
		return ErrStrLenMaxValidator
	}

	return nil
}

func NewStrLenMaxValidator() Validator {
	return &StrLenMaxValidator{}
}
