// Package raindrops provides a sound literal depending on the provided number is divisible by 3,4,7
package raindrops

import (
	"strconv"
)

// Convert function takes in an integer and return sound string
func Convert(i int) string {
	returnSound := ""
	if i%3 == 0 {
		returnSound += "Pling"
	}
	if i%5 == 0 {
		returnSound += "Plang"
	}
	if i%7 == 0 {
		returnSound += "Plong"
	}

	if returnSound == "" {
		returnSound = strconv.Itoa(i)
	}
	return returnSound

}
