package cipher

import (
	"strings"
	"unicode"
)

func rotate(r rune, shift int) rune {
	if r >= 'a' && r <= 'z' {
		return 'a' + (r-'a'+rune(shift)+26)%26
	}
	return r
}

func alphabetIndex(r rune) int {
	switch {
	case r >= 'a' && r <= 'z':
		return int(r - 'a')
	case r >= 'A' && r <= 'Z':
		return int(r - 'A')
	default:
		return -1
	}
}

type shift struct {
	distance int
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift{distance: distance}
}

func (s shift) Encode(input string) string {
	return shiftText(input, s.distance)
}

func (s shift) Decode(input string) string {
	return shiftText(input, -s.distance)
}

func shiftText(input string, distance int) string {
	var b strings.Builder
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			b.WriteRune(rotate(r, distance))
		}
	}
	return b.String()
}

type vigenere struct {
	key string
}

func NewVigenere(key string) Cipher {
	if !isValidVigenereKey(key) {
		return nil
	}
	return vigenere{key: key}
}

func isValidVigenereKey(key string) bool {
	if len(key) < 2 {
		return false
	}

	for _, r := range key {
		if r < 'a' || r > 'z' {
			return false
		}
	}

	for _, r := range key {
		if r != rune(key[0]) {
			return true
		}
	}
	return false
}

func (v vigenere) Encode(input string) string {
	return v.transform(input, 1)
}

func (v vigenere) Decode(input string) string {
	return v.transform(input, -1)
}

func (v vigenere) transform(input string, direction int) string {
	var b strings.Builder
	keyLen := len(v.key)
	keyIndex := 0

	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			shift := alphabetIndex(rune(v.key[keyIndex%keyLen])) * direction
			b.WriteRune(rotate(r, shift))
			keyIndex++
		}
	}
	return b.String()
}
