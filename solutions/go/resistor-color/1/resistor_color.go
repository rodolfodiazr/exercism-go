package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
        "black", "brown", "red",
        "orange", "yellow", "green",
        "blue", "violet", "grey",
        "white",
    }

}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    c := Colors()
	for i := 0; i < len(c); i++ {
        if c[i] == color {
            return i
        }
    }
    
    return 0
}
