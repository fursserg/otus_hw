package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/fixme_my_friend/hw09_struct_validator/validators" //nolint:depguard
)

var (
	ErrEmptyRule          = errors.New("empty rule")
	ErrUndefinedValidator = errors.New("unsupported validation")
	ErrWrongFormatRule    = errors.New("wrong format rule")
	ErrNotStruct          = errors.New("only structs can validate")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	str := strings.Builder{}

	for _, x := range v {
		str.WriteString(x.Field + ": " + fmt.Sprintf("%s", x.Err) + "\n")
	}

	return str.String()
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors
	var valErr ValidationErrors
	st := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	for i := 0; i < val.NumField(); i++ {
		field := st.Field(i)

		validate, ok := field.Tag.Lookup("validate")
		if !ok {
			continue
		}

		fieldValue := val.FieldByName(field.Name).Interface()

		if validate == "" {
			return fmt.Errorf("getting validator for field %s: %w", field.Name, ErrEmptyRule)
		}

		rules := strings.Split(validate, "|")

		for _, rule := range rules {
			err := validateField(val, field, rule, fieldValue)

			if err == nil {
				continue
			}

			if errors.As(err, &valErr) {
				validationErrors = append(validationErrors, err.(ValidationErrors)...) //nolint:errorlint
			} else {
				return err
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func getValidator(valueType string, validator string) (validators.Validator, error) {
	r := valueType + ":" + validator

	switch r {
	case "string:lenMin":
		return validators.NewStrLenMinValidator(), nil
	case "string:len":
		return validators.NewStrLenEqualValidator(), nil
	case "string:lenMax":
		return validators.NewStrLenMaxValidator(), nil
	case "string:regexp":
		return validators.NewStrRegexpValidator(), nil
	case "string:in":
		return validators.NewStrInValidator(), nil
	case "int:min":
		return validators.NewIntMinValidator(), nil
	case "int:max":
		return validators.NewIntMaxValidator(), nil
	case "int:range":
		return validators.NewIntRangeValidator(), nil
	case "int:in":
		return validators.NewIntInValidator(), nil
	default:
		return nil, ErrUndefinedValidator
	}
}

func validateField(val reflect.Value, field reflect.StructField, rule string, fieldValue any) error {
	var validationErrors ValidationErrors
	var pe validators.ParsingError

	limitationType, limitationValue, found := strings.Cut(rule, ":")
	if !found {
		return fmt.Errorf("getting validator for field %s: %w", field.Name, ErrWrongFormatRule)
	}

	isArray := val.FieldByName(field.Name).Kind() == reflect.Slice || val.FieldByName(field.Name).Kind() == reflect.Array
	length := 1
	fieldType := val.FieldByName(field.Name).Kind().String()

	if isArray {
		length = val.FieldByName(field.Name).Len()
		fieldType = reflect.TypeOf(fieldValue).Elem().String()
	}

	validator, err := getValidator(fieldType, limitationType)
	if err != nil {
		return fmt.Errorf("getting validator for field %s: %w", field.Name, err)
	}

	validator.Limitation(limitationValue)

	for i := 0; i < length; i++ {
		fieldName := field.Name

		validationValue := fieldValue

		if isArray {
			fieldName = fmt.Sprintf("%s[%d]", fieldName, i)
			validationValue = reflect.ValueOf(fieldValue).Index(i).Interface()
		}

		validator.Actual(validationValue)

		if err := validator.Validate(); err != nil {
			if errors.As(err, &pe) {
				return fmt.Errorf(`parsing field "%s": %w`, fieldName, err)
			}

			validationErrors = append(validationErrors, ValidationError{Field: fieldName, Err: err})
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
