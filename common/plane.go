package common

import (
	_ "fmt"
	"math"
	"math/rand"
)

type Plane struct {
	normal Vector
	BasicObject
}

func NewPlane() *Plane {
	return &Plane{
		BasicObject: BasicObject{
			transform: NewIdentity4(),
			inverseTransform: NewIdentity4(),
			inverseTransposeTransform: NewIdentity4(),
			material:  NewMaterial(),
			id:        rand.Int63(),
		},
	}
}

func (p *Plane) NormalAt(point Vector) Vector {
	if p.normal.Eq(New3Vector(0, 0, 0)) {
		normal := p.transformVectorToWorld(
			New3Vector(0, 1, 0),
		)

		normal.w = 0

		p.normal = normal
	}

	return p.normal
}

func (p *Plane) Intersect(r Ray) *Intersections {
	r = p.transformRayToObject(r)

	if math.Abs(r.direction.y) < Epsilon {
		return &Intersections{}
	}

	return &Intersections{
		[]*Intersection{
			&Intersection{
				p, -r.origin.y / r.direction.y,
			},
		},
	}
}
