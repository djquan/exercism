// Package space provides a method for determining age based on planet
package space

// Planet represents a Planet
type Planet string

const secondsInAnEarthYear = 365.24 * 24 * 60 * 60

// Age calculates the age in years based on seconds and planet
func Age(s float64, p Planet) float64 {
	return s / (p.orbitalPeriod() * secondsInAnEarthYear)
}

func (p Planet) orbitalPeriod() float64 {
	switch p {
	case "Earth":
		return 1
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	default:
		panic("Unknown planet")
	}
}
