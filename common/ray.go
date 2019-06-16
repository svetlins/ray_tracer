package common

type Ray struct {
	origin    Vector
	direction Vector
}

func NewRay(origin Vector, direction Vector) Ray {
	return Ray{origin, direction}
}

func (ray Ray) Position(t float64) Vector {
	return ray.origin.Add(ray.direction.Scale(t))
}

func (ray Ray) Transform(m Matrix) Ray {
	return Ray{
		m.Transform(ray.origin),
		m.Transform(ray.direction),
	}
}
