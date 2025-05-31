package validators

type Validator interface {
	Validate() error
	Limitation(string)
	Actual(any)
}

type AbstractValidator struct {
	limitation string
	actual     any
}

type ParsingError struct {
	message string
}

func (pe ParsingError) Error() string {
	return pe.message
}

func (a *AbstractValidator) Limitation(lim string) {
	a.limitation = lim
}

func (a *AbstractValidator) Actual(act any) {
	a.actual = act
}

func (a *AbstractValidator) Validate() error {
	return nil
}

func (a *AbstractValidator) parsingError(s string) error {
	return &ParsingError{s}
}
