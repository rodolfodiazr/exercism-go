package raindrops

import "strconv"

func Convert(number int) string {
    sounds := map[int]string{
        3: "Pling",
        5: "Plang",
        7: "Plong",
    }

    var result string
    for factor, sound := range sounds {
        if number%factor == 0 {
            result += sound
        }
    }

    if result == "" {
        return strconv.Itoa(number)
    }

    return result
}
