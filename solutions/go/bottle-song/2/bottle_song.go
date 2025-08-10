package bottlesong

import (
    "fmt"
    "strings"
)

var numbers = map[int]string{
    0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func plural(n int) string {
    if n == 1 {
        return ""
    }
    return "s"
}

func verse(n int) []string {
    next := n - 1
    return []string{
        fmt.Sprintf(
            "%s green bottle%s hanging on the wall,",
            numbers[n], 
            plural(n),
        ),
        fmt.Sprintf(
            "%s green bottle%s hanging on the wall,",
            numbers[n],
            plural(n),
        ),
        "And if one green bottle should accidentally fall,",
        fmt.Sprintf(
            "There'll be %s green bottle%s hanging on the wall.", 
            strings.ToLower(numbers[next]), plural(next),
        ),
    }
}

func Recite(startBottles, takeDown int) []string {
    var out []string
    end := startBottles - (takeDown - 1)
    
    for i := startBottles; i >= end; i-- {
        out = append(out, verse(i)...)
        if i > end {
            out = append(out, "")
        }
    }
	return out
}