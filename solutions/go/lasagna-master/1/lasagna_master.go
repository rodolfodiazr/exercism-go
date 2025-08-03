package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        timePerLayer = 2
    }
    return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var s, n int
    for _, l := range layers {
        if l == "sauce" {
            s++
        	continue   
        }

        if l == "noodles" {
            n++
        }
    }
    return (n * 50), (float64(s) * 0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(a []float64, numberOfPortions int) []float64 {
    amounts := []float64{}
    amounts = append(amounts, a...)
    for i := range a {
        amounts[i] = (amounts[i] / 2) * float64(numberOfPortions)
    }
    return amounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
