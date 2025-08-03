package isbn

import (
    "strings"
    "strconv"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
        return false
    }

    idx := 10
    sum := 0
    for _, c := range isbn {
        isLastCharacter := idx == 1
        if isLastCharacter && c == 'X' {
            sum += 10
            continue
        }
        
        digit, err := strconv.Atoi(string(c))
        if err != nil {
            return false
        }
        
        sum += (digit * (idx))
        idx--
    }
    
    return sum % 11 == 0
}
