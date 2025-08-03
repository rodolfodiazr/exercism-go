package resistorcolorduo

var BANDS = map[string]int{
    "black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4,
    "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) < 2 {
        return 0
	}

    first, ok := BANDS[colors[0]]
    if !ok {
        return 0
    }
        
    second, ok := BANDS[colors[1]]
    if !ok {
        return 0
    }

	return (first * 10) + second
}
