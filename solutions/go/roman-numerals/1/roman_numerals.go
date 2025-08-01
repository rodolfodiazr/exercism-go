package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("input must be between 1 and 3999")
	}

	return toRomanNumeral(input), nil
}

func toRomanNumeral(n int) string {
	numerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, numeral := range numerals {
		if n >= numeral.value {
			return numeral.symbol + toRomanNumeral(n-numeral.value)
		}
	}
	return ""
}
