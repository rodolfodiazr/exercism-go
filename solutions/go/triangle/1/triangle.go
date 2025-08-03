// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

type Kind int

const (
    NaT Kind = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

func isValid(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return a+b >= c && b+c >= a && a+c >= b
}

// KindFromSides returns the triangle kind given its three side lengths.
func KindFromSides(a, b, c float64) Kind {
	if !isValid(a, b, c) {
        return NaT
    }
    
    switch {
    case a == b && b == c: 
        return Equ
	case a == b || b == c || a == c: 
    	return Iso
    default: 
        return Sca
    }
}
