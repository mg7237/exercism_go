package luhn

import (
	"strconv"
	"strings"
)

// Valid function checks if a number sequence is valid Luhn
func Valid(input string) bool {
	counter, i := 0, 0
	input = strings.ReplaceAll(input, " ", "")
	inputLength := len(input)
	if inputLength < 2 {
		return false
	}
	for i = inputLength - 1; i >= 0; i-- {
		intAtPosition, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return false
		}
		if (inputLength-i)%2 == 0 {
			intAtPosition = intAtPosition * 2
			if intAtPosition > 9 {
				intAtPosition = intAtPosition - 9
			}
		}
		counter += intAtPosition
	}
	return counter%10 == 0
}
