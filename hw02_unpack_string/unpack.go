package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrTooManyChars  = errors.New("too many characters")
)

func Unpack(str string) (string, error) {
	var res strings.Builder
	runes := []rune(str)

	if valid, err := isValid(str); !valid || err != nil {
		return "", ErrInvalidString
	}

	for i := 0; i < len(runes); i++ {
		symbol := string(runes[i])

		next := ""
		if len(runes) > i+1 {
			next = string(runes[i+1])
		}

		prev := ""
		if i > 0 {
			prev = string(runes[i-1])
		}

		prevPrev := ""
		if i > 1 {
			prevPrev = string(runes[i-2])
		}

		isNextNum, nextNum, nextNumErr := isNumeric(next)
		if next != "" && nextNumErr != nil {
			return "", ErrInvalidString
		}

		if symbol == "\\" && (prev != "\\" || prevPrev == "\\") {
			continue
		}

		if isNextNum {
			res.WriteString(strings.Repeat(symbol, nextNum))
			i++
			continue
		}

		res.WriteString(symbol)
	}

	return res.String(), nil
}

func isNumeric(s string) (bool, int, error) {
	if len([]rune(s)) != 1 {
		return false, 0, ErrTooManyChars
	}

	if (s[0] > 47) && (s[0] < 58) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return false, 0, err
		}

		return true, num, nil
	}

	return false, 0, nil
}

func isValid(s string) (bool, error) {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		symbol := string(runes[i])

		next := ""
		if len(runes) > i+1 {
			next = string(runes[i+1])
		}

		prev := ""
		if i > 0 {
			prev = string(runes[i-1])
		}

		prevPrev := ""
		if i > 1 {
			prevPrev = string(runes[i-2])
		}

		isNum, _, curNumErr := isNumeric(symbol)
		if curNumErr != nil {
			return false, ErrInvalidString
		}

		isNextNum, _, nextNumErr := isNumeric(next)
		if next != "" && nextNumErr != nil {
			return false, ErrInvalidString
		}

		isPrevNum, _, prevNumErr := isNumeric(prev)
		if prev != "" && prevNumErr != nil {
			return false, ErrInvalidString
		}

		if symbol == "\\" && next == "" {
			return false, ErrInvalidString
		}

		if symbol == "\\" && !isNextNum && next != "\\" {
			return false, ErrInvalidString
		}

		if isNum && i == 0 {
			return false, ErrInvalidString
		}

		if isNum && isPrevNum && i > 1 && prevPrev != "\\" {
			return false, ErrInvalidString
		}
	}

	return true, nil
}
