package common

import (
	"math"
)

const Epsilon = 0.0001

type Vector struct {
	x float64
	y float64
	z float64
	w float64
}

func (t Vector) IsPoint() bool {
	return floatEqual(t.w, 1)
}

func NewPoint(x, y, z float64) Vector {
	return Vector{x, y, z, 1}
}

func New3Vector(x, y, z float64) Vector {
	return Vector{x, y, z, 0}
}

func NewColor(r, g, b float64) Vector {
	return Vector{r, g, b, 0}
}

func (a Vector) Eq(b Vector) bool {
	return floatEqual(a.x, b.x) &&
		floatEqual(a.y, b.y) &&
		floatEqual(a.z, b.z) &&
		floatEqual(a.w, b.w)
}

func (a Vector) NotEq(b Vector) bool {
	return !a.Eq(b)
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}

}

func (a Vector) Subtract(b Vector) Vector {
	return Vector{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}

}

func (a Vector) Negate() Vector {
	return Vector{-a.x, -a.y, -a.z, -a.w}
}

func (a Vector) Scale(factor float64) Vector {
	return Vector{a.x * factor, a.y * factor, a.z * factor, a.w * factor}
}

func (a Vector) Divide(factor float64) Vector {
	return a.Scale(1.0 / factor)
}

func (a Vector) Multiply(b Vector) Vector {
	return Vector{a.x * b.x, a.y * b.y, a.z * b.z, a.w * b.w}
}

func (a Vector) Magnitude() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

func (a Vector) Normalize() Vector {
	magnitude := a.Magnitude()
	return New3Vector(a.x/magnitude, a.y/magnitude, a.z/magnitude)
}

func (a Vector) Dot(b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w
}

func (a Vector) Cross(b Vector) Vector {
	return Vector{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
		0,
	}

}

func (a Vector) Reflect(n Vector) Vector {
	n = n.Normalize()
	return a.Subtract(n.Scale(2 * a.Dot(n)))
}

func floatEqual(a float64, b float64) bool {
	return math.Abs(a-b) < Epsilon
}
