package wordy

import (
    "strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.TrimSpace(question)
	if !strings.HasPrefix(question, "What is ") || !strings.HasSuffix(question, "?") {
		return 0, false
	}

	expression := strings.TrimSuffix(strings.TrimPrefix(question, "What is "), "?")
	expression = strings.ReplaceAll(expression, "multiplied by", "multiplied_by")
	expression = strings.ReplaceAll(expression, "divided by", "divided_by")

	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, false
	}

	current, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	for i := 1; i < len(tokens); i += 2 {
		if i+1 >= len(tokens) {
			return 0, false
		}

		op := tokens[i]
		nextNumStr := tokens[i+1]

		nextNum, err := strconv.Atoi(nextNumStr)
		if err != nil {
			return 0, false
		}

		switch op {
		case "plus":
			current += nextNum
		case "minus":
			current -= nextNum
		case "multiplied_by":
			current *= nextNum
		case "divided_by":
			if nextNum == 0 {
				return 0, false
			}
			current /= nextNum
		default:
			return 0, false
		}
	}

	return current, true
}
