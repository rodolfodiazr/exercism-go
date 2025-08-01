package logs

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	m := make(map[string]string)
	m["❗"] = "recommendation"
	m["🔍"] = "search"
	m["☀"] = "weather"
	for _, st := range log {
		if value, ok := m[fmt.Sprintf("%c", st)]; ok {
			return value
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(
        log, 
        fmt.Sprintf("%c", oldRune),
        fmt.Sprintf("%c", newRune),
    )
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
