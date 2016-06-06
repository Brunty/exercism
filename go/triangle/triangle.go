package triangle

// A Kind denotes the type of triangle
type Kind uint8

// Kinds of triangles we can possibly have
const (
	Equ Kind = iota // equilateral
	Iso             // isosceles
	Sca             // scalene
	NaT             // not a triangle
)

// KindFromSides returns the Kind of the triangle it is.
// it's defined by the lengths of its sides a, b and c.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !(a+b > c && a+c > b && b+c > a):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}
}
