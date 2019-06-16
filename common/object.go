package common

import (
	"math"
	"math/rand"
	"sort"
)

type Object interface {
	Intersect(r Ray) *Intersections
	SetTransform(m Matrix)
	SetMaterial(m Material)
	Material() Material
	NormalAt(p Vector) Vector
	Id() int64

	transformPointToObject(p Vector) Vector
}

type Sphere struct {
	BasicObject
}

type Intersection struct {
	object Object
	t      float64
}

type Intersections struct {
	intersections []*Intersection
}

func NewSphere() *Sphere {
	return &Sphere{
		BasicObject: BasicObject{
			material: NewMaterial(),
			transform: NewIdentity4(),
			inverseTransform: NewIdentity4(),
			inverseTransposeTransform: NewIdentity4(),
			id: rand.Int63(),
		},
	}
}

func NewGlassSphere() *Sphere {
	return NewSphere()

	// sphere := &Sphere{
	// 	BasicObject{
	// 		transform: NewIdentity4(),
	// 		material:  Glass(),
	// 		id:        rand.Int63(),
	// 	},
	// }

	// return sphere
}

func (s1 *Sphere) Eq(s2 *Sphere) bool {
	return s1.transform.Equal(s2.transform) &&
		s1.material.Equal(s2.material)
}

func (s *Sphere) NormalAt(p Vector) Vector {
	spherePosition := NewPoint(0, 0, 0)
	objectPoint := s.transformPointToObject(p)

	objectNormal := objectPoint.Subtract(spherePosition)

	worldNormal := s.transformVectorToWorld(objectNormal)
	worldNormal.w = 0

	return worldNormal.Normalize()
}

func (s *Sphere) Intersect(r Ray) *Intersections {
	spherePosition := NewPoint(0, 0, 0)

	r = s.transformRayToObject(r)

	sphereToRay := r.origin.Subtract(spherePosition)

	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return &Intersections{}
	} else {
		i1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		i2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		return NewIntersections(
			&Intersection{s, i1},
			&Intersection{s, i2},
		)

	}
}

func NewIntersections(is ...*Intersection) *Intersections {
	result := Intersections{}
	sort.Slice(is, func(i, j int) bool { return is[i].t < is[j].t })
	result.intersections = is

	return &result
}

func (is *Intersections) Len() int {
	return len(is.intersections)
}

func (is *Intersections) At(index int) *Intersection {
	return is.intersections[index]
}

func (is *Intersections) Hit() *Intersection {
	for _, potentialHit := range is.intersections {
		if potentialHit.t > 0 {
			return potentialHit
		}
	}

	return nil
}

func (is1 *Intersections) Merge(is2 *Intersections) {
	mergedIntersections := append(is1.intersections, is2.intersections...)

	sort.Slice(
		mergedIntersections,
		func(i, j int) bool {
			return mergedIntersections[i].t < mergedIntersections[j].t
		},
	)

	is1.intersections = mergedIntersections
}

func (i *Intersection) T() float64 {
	return i.t
}
