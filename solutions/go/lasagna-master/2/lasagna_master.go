package lasagna

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
    if avgPrepTimePerLayer == 0 {
        avgPrepTimePerLayer = 2
    }
    return len(layers) * avgPrepTimePerLayer
}

func Quantities(layers []string) (int, float64) {
    var amountOfSauce, amountOfNoodles int
    for _, layer := range layers {
        switch layer {
        case "sauce": 
            amountOfSauce++
        case "noodles":
            amountOfNoodles++
        }
    }
    return (amountOfNoodles * 50), (float64(amountOfSauce) * 0.2)
}

func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(amounts []float64, numberOfPortions int) []float64 {
    var result = make([]float64, len(amounts))
    for i := range amounts {
        result[i] = (amounts[i] * float64(numberOfPortions)) / 2
    }
    return result
}	