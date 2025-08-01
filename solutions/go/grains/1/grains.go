package grains

import "errors"


func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("input must be between 1 and 64")
	}
    
	return 1 << (number - 1), nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		total += sq
	}
	return total
}
