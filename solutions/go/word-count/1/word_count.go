package wordcount

import (
    "strings"
    "unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    splitFn := func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
    }

    var output Frequency = map[string]int{}
    words := strings.FieldsFunc(phrase, splitFn)
    for _, word := range words {
		word, _ = strings.CutPrefix(word, "'")
        word, _ = strings.CutSuffix(word, "'")
        
        if word == "" {
            continue
        }
        output[strings.ToLower(word)]++
    }
	return output
}