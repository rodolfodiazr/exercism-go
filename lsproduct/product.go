package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	for _, r := range digits {
		if r < '0' || r > '9' {
			return 0, errors.New("digits input must only contain digits")
		}
	}

	var max int64
	for i := 0; i <= len(digits)-span; i++ {
		var product int64 = 1
		for j := 0; j < span; j++ {
			digit := int64(digits[i+j] - '0')
			product *= digit
		}

		if product > max {
			max = product
		}
	}

	return max, nil
}
