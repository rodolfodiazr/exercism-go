// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "unicode"
)

// Hey returns Bob's response based on the tone and content of the input remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
        return "Fine. Be that way!"
	}

	if isQuestion(remark) && shouting(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion(remark) {
		return "Sure."
	}
    
    if shouting(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func shouting(input string) bool {
    letters := 0
    uppers := 0
    
    for _, c := range input {
		if unicode.IsLetter(c) {
			letters++
			if unicode.IsUpper(c) {
				uppers++
			}
        }
    }

	return letters > 0 && letters == uppers
} 

func isQuestion(input string) bool {
    return strings.HasSuffix(input, "?")
}
