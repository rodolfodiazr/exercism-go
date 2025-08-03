package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers
// Note: The loop can be replaced by the formula `n * (n + 1) / 2`
func SquareOfSum(n int) int {
    var sum int
	for i := 1; i <= n; i ++ {
        sum += i
    }
    return sum * sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
// Note: The loop can be removed, and the sum calculated 
// with the formula n * (n + 1) * (2*n + 1) / 6
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i ++ {
        sum += i * i
    }
    return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
