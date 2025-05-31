package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/fixme_my_friend/hw09_struct_validator/validators" //nolint:depguard
	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		title       string
		in          interface{}
		expectedErr error
	}{
		{
			"all fields are wrong",
			User{
				ID:     strings.Repeat("0", 1),
				Name:   "John",
				Age:    3,
				Email:  "wrong@mail",
				Role:   "guest",
				Phones: []string{""},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{Field: "ID", Err: validators.ErrStrLenEqualValidator},
				ValidationError{Field: "Age", Err: validators.ErrIntMinValidator},
				ValidationError{Field: "Email", Err: validators.ErrStrRegexpValidator},
				ValidationError{Field: "Role", Err: validators.ErrStrInlValidator},
				ValidationError{Field: "Phones[0]", Err: validators.ErrStrLenEqualValidator},
			},
		},
		{
			"Only ID is correct",
			User{
				ID:     strings.Repeat("0", 36),
				Name:   "John",
				Age:    99,
				Email:  "wrong@mail",
				Role:   "guest",
				Phones: []string{"123"},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{Field: "Age", Err: validators.ErrIntMaxValidator},
				ValidationError{Field: "Email", Err: validators.ErrStrRegexpValidator},
				ValidationError{Field: "Role", Err: validators.ErrStrInlValidator},
				ValidationError{Field: "Phones[0]", Err: validators.ErrStrLenEqualValidator},
			},
		},
		{
			"ID, Age are correct",
			User{
				ID:     strings.Repeat("0", 36),
				Name:   "John",
				Age:    20,
				Email:  "wrong@mail",
				Role:   "guest",
				Phones: []string{"", ""},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{Field: "Email", Err: validators.ErrStrRegexpValidator},
				ValidationError{Field: "Role", Err: validators.ErrStrInlValidator},
				ValidationError{Field: "Phones[0]", Err: validators.ErrStrLenEqualValidator},
				ValidationError{Field: "Phones[1]", Err: validators.ErrStrLenEqualValidator},
			},
		},
		{
			"ID, Age, Email are correct",
			User{
				ID:     strings.Repeat("0", 36),
				Name:   "John",
				Age:    20,
				Email:  "correct@email.test",
				Role:   "guest",
				Phones: []string{strings.Repeat("0", 11), ""},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{Field: "Role", Err: validators.ErrStrInlValidator},
				ValidationError{Field: "Phones[1]", Err: validators.ErrStrLenEqualValidator},
			},
		},
		{
			"ID, Age, Email, Role are correct",
			User{
				ID:     strings.Repeat("0", 36),
				Name:   "John",
				Age:    20,
				Email:  "correct@email.test",
				Role:   "admin",
				Phones: []string{strings.Repeat("0", 11), ""},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{Field: "Phones[1]", Err: validators.ErrStrLenEqualValidator},
			},
		},
		{
			"No errors",
			User{
				ID:     strings.Repeat("0", 36),
				Name:   "John",
				Age:    20,
				Email:  "correct@email.test",
				Role:   "admin",
				Phones: []string{strings.Repeat("0", 11), strings.Repeat("0", 11)},
				meta:   nil,
			},
			nil,
		},
		{
			"empty string",
			App{Version: ""},
			ValidationErrors{
				ValidationError{Field: "Version", Err: validators.ErrStrLenEqualValidator},
			},
		},
		{
			"struct without validation tags",
			Token{Header: []byte("header"), Payload: []byte("payload"), Signature: []byte("")},
			nil,
		},
		{
			"int not in set",
			Response{Code: 302, Body: "Moved temporarily"},
			ValidationErrors{
				ValidationError{Field: "Code", Err: validators.ErrIntInlValidator},
			},
		},
		{
			"int in set, no error",
			Response{Code: 200, Body: "ok"},
			nil,
		},
		{
			"wrong datatype",
			[]string{},
			ErrNotStruct,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case %s", tt.title), func(t *testing.T) {
			tt := tt
			t.Parallel()

			result := Validate(tt.in)
			require.Equal(t, tt.expectedErr, result)

			_ = tt
		})
	}
}
