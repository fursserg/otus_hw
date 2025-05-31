package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var ErrIntRangeValidator = errors.New("is not in range")

type IntRangeValidator struct {
	AbstractValidator
}

func (s *IntRangeValidator) Validate() error {
	fromTo := strings.Split(s.limitation, ",")
	if len(fromTo) != 2 {
		return s.parsingError(fmt.Sprintf("range validator accept only two arguments, passed %d", len(fromTo)))
	}

	from, errFrom := strconv.Atoi(fromTo[0])
	if errFrom != nil {
		return s.parsingError("range first argument has invalid format")
	}

	to, errTo := strconv.Atoi(fromTo[1])
	if errTo != nil {
		return s.parsingError("range second argument has invalid format")
	}

	actual := int(reflect.ValueOf(s.actual).Int())

	if actual < from || actual > to {
		return ErrIntRangeValidator
	}

	return nil
}

func NewIntRangeValidator() Validator {
	return &IntRangeValidator{}
}
