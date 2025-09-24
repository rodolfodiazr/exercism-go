package summultiples

func SumMultiples(limit int, divisors ...int) int {
    if len(divisors) == 0 {
        return 0
    }
    
    multiples := make(map[int]struct{})
    for _, divisor := range divisors {
        if divisor <= 0 {
            continue
        }
    	for i := divisor; i < limit; i += divisor {
            multiples[i] = struct{}{}
    	}
    }

    sum := 0
    for k := range multiples {
        sum += k
    }
    return sum
}


