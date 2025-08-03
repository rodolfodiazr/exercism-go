// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear returns true if the given year is a leap year.
// A year is a leap year if it is divisible by 4,
// except for years that are divisible by 100,
// unless they are also divisible by 400.
func IsLeapYear(year int) bool {
	if year % 100 == 0 {
		return year % 400 == 0
	}
	return year % 4 == 0
}
