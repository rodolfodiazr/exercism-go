package pangram

import "strings"

func IsPangram(input string) bool {
    input = strings.ToUpper(input)
    letters := map[rune]bool{}
	for _, r := range input {
        if r >= 'A' && r <= 'Z' {
            letters[r] = true
        }

        if len(letters) == 26 {
            return true
        }
    }
    return len(letters) == 26
}
