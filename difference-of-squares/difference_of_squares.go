package diffsquares

// SquareOfSum returns square of the sum of digits from 0 - input number
func SquareOfSum(input int) int {
	sumOfNumbers := 0
	for i := 1; i <= input; i++ {
		sumOfNumbers += i
	}
	return sumOfNumbers * sumOfNumbers
}

// SumOfSquares returns sum of the square of digits from 0 - input number
func SumOfSquares(input int) int {
	sumOfSquare := 0
	for i := 1; i <= input; i++ {
		sumOfSquare += i * i
	}
	return sumOfSquare
}

// Difference returns dufference of SquareOfSum and SumOfSquare for any input int
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
