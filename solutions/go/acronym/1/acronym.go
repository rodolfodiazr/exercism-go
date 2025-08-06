// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
    "strings"
    "unicode"
)

// Abbreviate returns the acronym formed 
// by taking the first letter of each word in the given string.
func Abbreviate(s string) string {
    s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	var b strings.Builder
	words := strings.Fields(s)

	for _, word := range words {
		for i, r := range word {
			if i == 0 {
				b.WriteRune(unicode.ToUpper(r))
				break
			}
		}
	}

	return b.String()
}
