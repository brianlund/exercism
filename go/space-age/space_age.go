// Package space provides a function for calculating age in different planet-years
package space

// Planet holds structure for a planet
type Planet = string

const earthSeconds float64 = 31557600

var spaceAge = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns corresponding years for different planets given an age in seconds
func Age(seconds float64, planet string) (age float64) {

	return seconds / (earthSeconds * spaceAge[planet])

}
