package isogram

import (
    "strings"
)

func IsIsogram(word string) bool {
    seen := make(map[rune]bool)
    word = strings.ToUpper(word)

    for _, c := range word {
        if c == '-' || c == ' ' {
            continue
        }

        if seen[c] {
            return false
        }
        
        seen[c] = true
    }

    return true
}
