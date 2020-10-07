// Package space calculates the input parameter's equivalent earth age using
// constants representing planet age in comparison to earth years
package space

import (
	"os"
)

// Age calculates the planet age in seconds
func Age(secondsToConvert float64, planet string) float64 {
	planetFigures := map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	value, ok := planetFigures[planet]
	if ok {
		age := secondsToConvert / (value * 365 * 24 * 60 * 60)
		println(planet, "Age: ", age)
		return age
	}
	println(planet, "ERROR: planet parameter not gound in planetFigures map ")
	os.Exit(1)
	return 0
}
