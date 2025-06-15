package validators

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

var ErrStrRegexpValidator = errors.New("does not match regular expression")

type StrRegexpValidator struct {
	AbstractValidator
}

func (s *StrRegexpValidator) Validate() error {
	re, err := regexp.Compile(s.limitation)
	if err != nil {
		return s.parsingError(fmt.Sprintf("compile regexp: %s", err))
	}

	actual := reflect.ValueOf(s.actual).String()

	if !re.MatchString(actual) {
		return ErrStrRegexpValidator
	}

	return nil
}

func NewStrRegexpValidator() Validator {
	return &StrRegexpValidator{}
}
