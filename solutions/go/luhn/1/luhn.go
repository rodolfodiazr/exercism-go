package luhn

import (
    "strconv"
    "strings"
)

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if len(id) <= 1 {
        return false
    }

	sum := 0
	shouldDouble := false
	for i := len(id) - 1; i >= 0; i-- {
    	num, err := strconv.Atoi(string(id[i]))
        if err != nil {
            return false
        }
        
        if shouldDouble {
            num *= 2
            if num > 9 {
                num -= 9
            }
        }

        sum += num
        shouldDouble = !shouldDouble
    }

    return sum % 10 == 0
}
