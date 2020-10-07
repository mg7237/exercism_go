package isogram

import (
	"strings"
)

// IsIsogram function decides if a string is Isogram or not.
// Returns bool value
func IsIsogram(inputString string) bool {
	trackerMap := map[byte]int{}
	isogram := true
	stringLength := len(inputString)
	ucaseInput := strings.ToUpper(inputString)
	var i, j int
	for i = 0; i < stringLength; i++ {
		j = trackerMap[ucaseInput[i]]
		if j != 0 && ucaseInput[i] != ' ' && ucaseInput[i] != '-' {
			isogram = false
			i = stringLength - 1
		} else {
			trackerMap[ucaseInput[i]] = '1'
		}
	}
	return isogram
}
