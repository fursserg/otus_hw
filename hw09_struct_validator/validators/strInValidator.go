package validators

import (
	"errors"
	"reflect"
	"slices"
	"strings"
)

var ErrStrInlValidator = errors.New("is not included in the set")

type StrInValidator struct {
	AbstractValidator
}

func (s *StrInValidator) Validate() error {
	in := strings.Split(s.limitation, ",")
	actual := reflect.ValueOf(s.actual).String()

	if !slices.Contains(in, actual) {
		return ErrStrInlValidator
	}

	return nil
}

func NewStrInValidator() Validator {
	return &StrInValidator{}
}
