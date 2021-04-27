// Package triangle implements functionality to define triangles
package triangle

import "math"

type Kind string

const (
	NaT = "Nat" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides will return what kind of triangle is passed on (Equ, Iso, Sca, Nat)
func KindFromSides(a, b, c float64) (k Kind) {
	isValid := func(s float64) bool {
		return math.IsInf(s, 0) || s <= 0
	}

	switch {
	case isValid(a), isValid(b), isValid(c),
		!(a+b >= c), !(b+c >= a), !(a+c >= b):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	case a != b && a != c && b != c:
		k = Sca
	}
	return k
}
