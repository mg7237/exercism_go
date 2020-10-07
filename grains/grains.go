package grains

import (
	"errors"
)

// Square return sum of squares of given input between 1 to 64
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("Invalid Input")
	}
	return 1 << (input - 1), nil
}

//Total calculates total grains on chess board
func Total() uint64 {
	return 1<<64 - 1
}
