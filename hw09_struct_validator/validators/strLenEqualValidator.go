package validators

import (
	"errors"
	"reflect"
	"strconv"
)

var ErrStrLenEqualValidator = errors.New("not equal to limitation")

type StrLenEqualValidator struct {
	AbstractValidator
}

func (s *StrLenEqualValidator) Validate() error {
	limitation, err := strconv.Atoi(s.limitation)

	actual := len(reflect.ValueOf(s.actual).String())

	if err != nil {
		return s.parsingError("invalid format")
	}

	if limitation > 0 && actual != limitation {
		return ErrStrLenEqualValidator
	}

	return nil
}

func NewStrLenEqualValidator() *StrLenEqualValidator {
	return &StrLenEqualValidator{}
}
