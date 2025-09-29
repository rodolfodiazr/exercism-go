package series

func All(n int, s string) []string {
    if n <= 0 || n > len(s) {
		return []string{}
	}
	var output []string
    for i := 0; i <= len(s)-n; i++ {
        output = append(output, s[i:i+n])
	}
    return output
}

func UnsafeFirst(n int, s string) string {
    if n <= 0 {
		return ""
	}
	if len(s) < n {
        return s
	}
    return s[:n]
}
