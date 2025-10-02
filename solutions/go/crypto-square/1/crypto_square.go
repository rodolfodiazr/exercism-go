package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func Encode(pt string) string {
	// Step 1: normalize
	var normalized strings.Builder
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normalized.WriteRune(unicode.ToLower(r))
		}
	}
	text := normalized.String()
	if len(text) == 0 {
		return ""
	}

	// Step 2: find rows and cols
	L := len(text)
	r := int(math.Floor(math.Sqrt(float64(L))))
	c := int(math.Ceil(math.Sqrt(float64(L))))
	if r*c < L {
		r++
	}

	// Step 3: build rectangle rows (pad with spaces)
	rows := make([]string, r)
	for i := 0; i < r; i++ {
		start := i * c
		end := start + c
		if end > L {
			end = L
		}
		row := text[start:end]
		if len(row) < c {
			row += strings.Repeat(" ", c-len(row))
		}
		rows[i] = row
	}

	// Step 4: read column-wise
	var encoded []string
	for col := 0; col < c; col++ {
		var chunk strings.Builder
		for row := 0; row < r; row++ {
			chunk.WriteByte(rows[row][col])
		}
		encoded = append(encoded, chunk.String())
	}

	return strings.Join(encoded, " ")
}
