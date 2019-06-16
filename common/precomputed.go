package common

type Precomputed struct {
	T                    float64
	Object               Object
	Point                Vector
	EyeVector            Vector
	NormalVector         Vector
	ReflectVector        Vector
	RefractiveIndexEnter float64
	RefractiveIndexExit  float64
	Inside               bool
	OverPoint            Vector
	UnderPoint           Vector
}

func PrepareComputations(hit *Intersection, r Ray, is *Intersections) *Precomputed {
	point := r.Position(hit.t)
	eyeVector := r.direction.Scale(-1)
	normalVector := hit.object.NormalAt(point)
	reflectVector := r.direction.Reflect(normalVector)
	inside := eyeVector.Dot(normalVector) < 0

	if inside {
		normalVector = normalVector.Scale(-1)
	}

	container := make([]Object, 0, 16)

	var refractiveIndexEnter, refractiveIndexExit float64

	for j := 0; j < is.Len(); j++ {
		intersection := is.At(j)

		if intersection.t == hit.t {
			if len(container) == 0 {
				refractiveIndexEnter = 1.0
			} else {
				refractiveIndexEnter =
					container[len(container) - 1].Material().refractiveIndex
			}
		}

		include := false
		var includeIndex = -1

		for i, object := range container {
			if object == intersection.object {
				include = true
				includeIndex = i
			}

		}

		if include {
			container = append(
				container[:includeIndex],
				container[includeIndex+1:]...,
			)
		} else {
			container = append(container, intersection.object)
		}


		if intersection.t == hit.t {
			if len(container) == 0 {
				refractiveIndexExit = 1.0
			} else {
				refractiveIndexExit = container[len(container) - 1].Material().refractiveIndex
			}
		}
	}

	return &Precomputed{
		T:             hit.t,
		Object:        hit.object,
		Point:         point,
		EyeVector:     eyeVector,
		NormalVector:  normalVector,
		ReflectVector: reflectVector,
		RefractiveIndexEnter: refractiveIndexEnter,
		RefractiveIndexExit: refractiveIndexExit,
		Inside:        inside,
		OverPoint:     point.Add(normalVector.Scale(Epsilon)),
		UnderPoint:     point.Subtract(normalVector.Scale(Epsilon)),
	}
}
