package anagram

import "strings"

func letterFrequency(str string) map[rune]int {
    frequency := make(map[rune]int)
    for _, c := range str {
        frequency[c]++
    }
    return frequency
}

func mapsEqual(a, b map[rune]int) bool {
    if len(a) != len(b) {
        return false
    }

    for k, _ := range a {
        if a[k] != b[k] {
            return false
        }
    }
    return true
}

func Detect(subject string, candidates []string) []string {
    output := []string{}
    subjectLower := strings.ToLower(subject)
    subjectMap := letterFrequency(subjectLower)

    for _, candidate := range candidates {
        candidateLower := strings.ToLower(candidate)
		if candidateLower == subjectLower || len(candidate) != len(subject) {
            continue
        }

        candidateMap := letterFrequency(candidateLower)
        if !mapsEqual(subjectMap, candidateMap) {
            continue
        }

        output = append(output, candidate)
    }

    return output
}
