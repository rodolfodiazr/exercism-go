package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    m := make(map[string]int)
	for score, letters := range in {
        for _, letter := range letters {
            m[strings.ToLower(letter)] = score
        }
    }
    return m
}
