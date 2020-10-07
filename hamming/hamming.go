// Package hamming calculates the hamming distances between two DNA strands
// provided as parameters to the Distance function within this package
package hamming

import "errors"

//Distance function follows:
// - Check if length of 2 strings provided is same. If not then return error
// - Loop through the byte Slice and check if byte character at each position
//   across Slice A and B should be equal. If not the increment difference counter
// return final counter value and nil error.
func Distance(a, b string) (int, error) {
	byteA := []byte(a)
	byteB := []byte(b)

	if len(byteA) != len(byteB) {
		return 0, errors.New("Error: DNA parameters lengths are not equal")
	}
	// Distance counter
	counter := 0

	for i := 0; i < len(byteA); i++ {
		if byteA[i] != byteB[i] {
			counter++
		}
	}
	return counter, nil
}
