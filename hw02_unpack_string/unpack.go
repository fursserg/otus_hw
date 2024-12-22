package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func main() {
	testStrings := []string{
		`a4bc2d5e`,
		"abcd",
		"3abc",
		"45",
		`aaa10b`,
		"aaa0b",
		"d0f4",
		"d10f4",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
		`qw\ne`,
		`qwe\\\3`,
	}

	for _, str := range testStrings {
		res, err := Unpack(str)
		if err != nil {
			res = fmt.Sprintf("%s", err)
		}

		fmt.Printf("%s => %s\n", str, res)
	}
}

func Unpack(str string) (string, error) {
	var res []string
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		symbol := string(runes[i])

		next := ""
		if len(runes) > i+1 {
			next = string(runes[i+1])
		}

		isNum, num := isNumeric(symbol)
		isNextNum, _ := isNumeric(next)

		if symbol == "\\" {
			if next == "" {
				return "", ErrInvalidString
			}

			if !isNextNum && next != "\\" {
				return "", ErrInvalidString
			}

			res = append(res, next)
			i++
			continue
		}

		if isNum {
			if i == 0 || isNextNum {
				return "", ErrInvalidString
			}

			res = append(res[:len(res)-1], strings.Split(strings.Repeat(res[len(res)-1], num), "")...)
			continue
		}

		res = append(res, symbol)
	}

	return strings.Join(res, ""), nil
}

func isNumeric(s string) (bool, int) {
	num, err := strconv.Atoi(s)
	if err == nil {
		return true, num
	}

	return false, 0
}
