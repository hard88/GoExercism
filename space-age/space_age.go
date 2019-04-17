package space

import (
	"strings"
)

type Planet string

func Age(second float64, planet Planet) float64 {
	earthYear := second/60/60/24/365.25
	var proportion float64
	if strings.Compare("Earth", string(planet)) == 0 {
		proportion = 1
	} else if strings.Compare("Mercury", string(planet)) == 0 {
		proportion = 0.2408467
	} else if strings.Compare("Venus", string(planet)) == 0 {
		proportion = 0.61519726
	} else if strings.Compare("Mars", string(planet)) == 0 {
		proportion = 1.8808158
	} else if strings.Compare("Jupiter", string(planet)) == 0 {
		proportion = 11.862615
	} else if strings.Compare("Saturn", string(planet)) == 0 {
		proportion = 29.447498
	} else if strings.Compare("Uranus", string(planet)) == 0 {
		proportion = 84.016846
	} else if strings.Compare("Neptune", string(planet)) == 0 {
		proportion = 164.79132
	}
	return earthYear/proportion
}
