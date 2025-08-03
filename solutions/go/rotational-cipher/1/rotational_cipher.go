package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
    shiftKey = shiftKey % 26
	var result []rune

	for _, r := range plain {
        switch {
        case unicode.IsUpper(r):
            rotated := 'A' + (r-'A'+rune(shiftKey)) % 26
            result = append(result, rotated)
        case unicode.IsLower(r):
        	rotated := 'a' + (r-'a'+rune(shiftKey)) % 26
        	result = append(result, rotated)
        default:
            result = append(result, r)
		}
	}
	return string(result)
}
