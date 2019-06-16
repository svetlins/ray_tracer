package common

import (
	"math"
)

func Transform(transforms ...Matrix) Matrix {
	result := NewIdentity4()

	for i := range transforms {
		result = result.Multiply(transforms[len(transforms) - 1 - i])
	}

	return result
}

func Translation(x, y, z float64) Matrix {
	return NewMatrix4([]float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	})
}

func Scaling(x, y, z float64) Matrix {
	return NewMatrix4([]float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
}

func RotationX(angle float64) Matrix {
	return NewMatrix4([]float64{
		1, 0, 0, 0,
		0, math.Cos(angle), -math.Sin(angle), 0,
		0, math.Sin(angle), math.Cos(angle), 0,
		0, 0, 0, 1,
	})
}

func RotationY(angle float64) Matrix {
	return NewMatrix4([]float64{
		math.Cos(angle), 0, math.Sin(angle), 0,
		0, 1, 0, 0,
		-math.Sin(angle), 0, math.Cos(angle), 0,
		0, 0, 0, 1,
	})
}

func RotationZ(angle float64) Matrix {
	return NewMatrix4([]float64{
		math.Cos(angle), -math.Sin(angle), 0, 0,
		math.Sin(angle), math.Cos(angle), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return NewMatrix4([]float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	})
}

func ViewTransform(from, to, up Vector) Matrix {
	eyeVector := to.Subtract(from).Normalize()
	left := eyeVector.Cross(up.Normalize())
	trueUp := left.Cross(eyeVector)
	orientation := NewMatrix4([]float64{
		left.x, left.y, left.z, 0,
		trueUp.x, trueUp.y, trueUp.z, 0,
		-eyeVector.x, -eyeVector.y, -eyeVector.z, 0,
		0, 0, 0, 1,
	})

	return orientation.Multiply(Translation(-from.x, -from.y, -from.z))
}
