package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		symbol := string(runes[i])

		prev := ""
		if i > 0 {
			prev = string(runes[i-1])
		}

		prevPrev := ""
		if i > 1 {
			prevPrev = string(runes[i-2])
		}

		// строка заканчивается на слэш
		if symbol == "\\" && len(runes)-1 == i {
			return "", ErrInvalidString
		}

		// слэш используется не только для экранирования слэшей и цифр
		if symbol == "\\" && !isNumeric(runes[i+1]) && string(runes[i+1]) != "\\" {
			return "", ErrInvalidString
		}

		// строка начианется с инта
		if isNumeric(runes[i]) && i == 0 {
			return "", ErrInvalidString
		}

		// проверка на два инта подряд
		if i > 1 && isNumeric(runes[i]) && isNumeric(runes[i-1]) && string(runes[i-2]) != "\\" {
			return "", ErrInvalidString
		}

		if symbol == "\\" && (prev != "\\" || prevPrev == "\\") {
			continue
		}

		if len(runes) > i+1 && isNumeric(runes[i+1]) {
			nextNum, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", ErrInvalidString
			}

			res.WriteString(strings.Repeat(symbol, nextNum))
			i++
			continue
		}

		res.WriteString(symbol)
	}

	return res.String(), nil
}

func isNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}
