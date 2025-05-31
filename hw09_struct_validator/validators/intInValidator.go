package validators

import (
	"errors"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

var ErrIntInlValidator = errors.New("is not included in the set")

type IntInValidator struct {
	AbstractValidator
}

func (s *IntInValidator) Validate() error {
	in := strings.Split(s.limitation, ",")
	actual := strconv.Itoa(int(reflect.ValueOf(s.actual).Int()))

	if !slices.Contains(in, actual) {
		return ErrIntInlValidator
	}

	return nil
}

func NewIntInValidator() Validator {
	return &IntInValidator{}
}
