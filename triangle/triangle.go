// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"go/constant"
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind constant.Kind

const (
    // Pick values for the following identifiers used by the test program.
    NaT = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	var k Kind

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) ||
		a <= 0 || b <= 0 || c <= 0 ||
		a + b < c || a + c < b || b + c < a { // is a triangle
		k = Kind(NaT)
	} else {
		if a == b && b == c {
			k = Kind(Equ)
		} else if a == b || b == c || a == c {
			k = Kind(Iso)
		} else {
			k = Kind(Sca)
		}
	}
	return k
}
