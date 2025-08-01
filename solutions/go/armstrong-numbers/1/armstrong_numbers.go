package armstrong

import "strconv"

func IsNumber(n int) bool {
	numberStr := strconv.Itoa(n)
	length := len(numberStr)
	sum := 0

	for _, c := range numberStr {
		digit := int(c - '0')
		sum += intPow(digit, length)
	}

	return sum == n
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}