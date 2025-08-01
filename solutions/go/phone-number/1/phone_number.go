package phonenumber

import (
    "errors"
    "fmt"
    "unicode"
    "strings"
)

func Number(phoneNumber string) (string, error) {
    var builder strings.Builder
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
            continue
		}
	}

	digits := builder.String()

	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", errors.New("11-digit numbers must start with '1'")
		}
		digits = digits[1:] // Clear the country code
	}

	if len(digits) != 10 {
		return "", errors.New("invalid number length; must be 10 digits")
	}

	if areaCode := digits[0]; areaCode < '2' || areaCode > '9' {
		return "", errors.New("area code must be a digit between 2 and 9")
	}
    
	if exchangeCode := digits[3]; exchangeCode < '2' || exchangeCode > '9' {
		return "", errors.New("exchange code must be a digit between 2 and 9")
	}
	return digits, nil
}

func AreaCode(phoneNumber string) (string, error) {
    digits, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
	return digits[0:3], nil
}

func Format(phoneNumber string) (string, error) {
    digits, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
	return fmt.Sprintf("(%v) %v-%v", digits[0:3], digits[3:6], digits[6:]), nil
}
