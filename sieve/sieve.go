package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	isComposite := make([]bool, limit+1)

	for p := 2; p*p <= limit; p++ {
		if !isComposite[p] {
			for multiple := p * p; multiple <= limit; multiple += p {
				isComposite[multiple] = true
			}
		}
	}

	var primes []int
	for i := 2; i <= limit; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
