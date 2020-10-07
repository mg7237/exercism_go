package scrabble

import "unicode"

// Score calculates the scrabble score for the string passed as parameter
func Score(testInput string) int {

	score := 0
	for _, char := range testInput {
		score += charValue(char)
	}
	return score
}

func charValue(char rune) int {
	switch unicode.ToUpper(char) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
