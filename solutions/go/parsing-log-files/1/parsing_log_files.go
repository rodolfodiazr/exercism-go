package parsinglogfiles

import (
    "fmt"
    "regexp"
)

func IsValidLine(text string) bool {
    pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
    re := regexp.MustCompile(pattern)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
  	pattern := `<[~*=\-]*>`
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
   	pattern := `\".*(?i)password.*\"`
    re := regexp.MustCompile(pattern)
    count := 0
    for _, line := range lines {
        if re.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line[0-9]+`
    re := regexp.MustCompile(pattern)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\w+)`
	re := regexp.MustCompile(pattern)
	results := []string{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			username := matches[1]
			results = append(results, fmt.Sprintf("[USR] %s %s", username, line))
            continue
		}
		
        results = append(results, line)
	}
	return results
}
