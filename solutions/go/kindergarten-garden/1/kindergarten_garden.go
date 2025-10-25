package kindergarten

import (
	"errors"
	"sort"
)

type Garden struct {
	rows     []string
	children []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	// Diagram must start with a newline
	if len(diagram) == 0 || diagram[0] != '\n' {
		return nil, errors.New("diagram must start with newline")
	}

	// Split into lines (ignore first empty one)
	lines := []string{}
	for _, line := range splitLines(diagram) {
		if line != "" {
			lines = append(lines, line)
		}
	}

	// Must have exactly two rows
	if len(lines) != 2 {
		return nil, errors.New("diagram must have exactly two rows")
	}

	// Rows must have equal even length
	if len(lines[0]) != len(lines[1]) || len(lines[0])%2 != 0 {
		return nil, errors.New("rows must have equal even length")
	}

	// Check for invalid characters
	valid := map[rune]bool{'G': true, 'C': true, 'R': true, 'V': true}
	for _, row := range lines {
		for _, r := range row {
			if !valid[r] {
				return nil, errors.New("invalid plant code in diagram")
			}
		}
	}

	// Check for duplicate children
	seen := map[string]bool{}
	for _, name := range children {
		if seen[name] {
			return nil, errors.New("duplicate child name")
		}
		seen[name] = true
	}

	// Check number of cups matches number of children
	if len(lines[0])/2 != len(children) {
		return nil, errors.New("cup count does not match number of children")
	}

	// Sort children alphabetically
	sorted := make([]string, len(children))
	copy(sorted, children)
	sort.Strings(sorted)

	return &Garden{
		rows:     lines,
		children: sorted,
	}, nil
}

func splitLines(s string) []string {
	lines := []string{}
	line := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, line)
			line = ""
		} else {
			line += string(r)
		}
	}
	lines = append(lines, line)
	return lines
}

func (g *Garden) Plants(child string) ([]string, bool) {
	index := -1
	for i, name := range g.children {
		if name == child {
			index = i
			break
		}
	}
    
	if index == -1 {
		return nil, false
	}

	start := index * 2
	end := start + 2

	plants := []rune(g.rows[0][start:end] + g.rows[1][start:end])

	mapping := map[rune]string{
		'G': "grass",
		'C': "clover",
		'R': "radishes",
		'V': "violets",
	}

	result := []string{}
	for _, p := range plants {
		result = append(result, mapping[p])
	}
	return result, true
}