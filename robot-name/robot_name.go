package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot type
type Robot string

var namesInUse = make(map[Robot]bool)

const maxN = 26 * 26 * 10 * 10 * 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Name receiver which gets name of existing robot or generates new Robot
func (r *Robot) Name() (string, error) {
	if string(*r) == "" {
		if len(namesInUse) == maxN {
			return "", errors.New("out of names")
		}

		var s Robot
		for {
			s = Robot(generate())
			if _, ok := namesInUse[s]; !ok {
				break
			}
		}
		namesInUse[s] = true
		*r = s
	}
	return string(*r), nil
}

func generate() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

//Reset the Robot name
func (r *Robot) Reset() {
	*r = Robot("")
}
