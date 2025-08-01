package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var list []rune
	for _, c := range s {
        if unicode.IsNumber(c) {
        	list = append(list, c)    
			continue
		}
        
		if unicode.IsLetter(c) {
            c = 'z' - (unicode.ToLower(c) - 'a')
			list = append(list, c)
            continue
		}
	}

	var builder strings.Builder
	for i, c := range list {
		if i > 0 && i % 5 == 0 {
			builder.WriteRune(' ')
		}
		builder.WriteRune(c)
	}

	return builder.String()
}