package common

import (
	"math"
	"testing"
)

func TestPrecomputed(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))
	sphere := NewSphere()
	is := NewIntersections(&Intersection{sphere, 4})
	comps := PrepareComputations(is.Hit(), ray, is)

	if comps.T != is.Hit().t {
		t.Errorf("Precomputed not working")
	}

	if comps.Object != sphere {
		t.Errorf("Precomputed not working")
	}

	if comps.Point.NotEq(NewPoint(0, 0, -1)) {
		t.Errorf("Precomputed not working")
	}

	if comps.EyeVector.NotEq(New3Vector(0, 0, -1)) {
		t.Errorf("Precomputed not working")
	}

	if comps.NormalVector.NotEq(New3Vector(0, 0, -1)) {
		t.Errorf("Precomputed not working")
	}

	if comps.Inside {
		t.Errorf("Precomputed not working")
	}
}

func TestPrecomputedInside(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 0), New3Vector(0, 0, 1))
	sphere := NewSphere()
	is := NewIntersections(&Intersection{sphere, 1})
	comps := PrepareComputations(is.Hit(), ray, is)

	if !comps.Inside {
		t.Errorf("Precomputed not working")
	}

	if comps.NormalVector.NotEq(New3Vector(0, 0, -1)) {
		t.Errorf("Precomputed not working")
	}
}

func TestPrecomputedOverPoint(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))
	sphere := NewSphere()
	sphere.SetTransform(Translation(0, 0, 1))
	intersections := NewIntersections(&Intersection{sphere, 5})
	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if precomputed.OverPoint.z >= -Epsilon/2 {
		t.Errorf("Precomputed Over Point not working")
	}

	if precomputed.Point.z < precomputed.OverPoint.z {
		t.Errorf("Precomputed Over Point not working")
	}
}

func TestPrecomputedReflected(t *testing.T) {
	ray := NewRay(
		NewPoint(0, 1, -1),
		New3Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2),
	)

	plane := NewPlane()

	intersections := NewIntersections(&Intersection{plane, math.Sqrt(2) / 2})
	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if precomputed.ReflectVector.NotEq(New3Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2)) {
		t.Errorf("Precomputed Over Point not working")
	}
}

func TestPrecomputedDifractiveIndices(t *testing.T) {
	sphere1 := NewGlassSphere()
	sphere1.SetTransform(Scaling(2, 2, 2))

	sphere2 := NewGlassSphere()
	sphere2.SetTransform(Translation(0, 0, -0.25))
	sphere2.material.refractiveIndex = 2.0

	sphere3 := NewGlassSphere()
	sphere3.SetTransform(Translation(0, 0, 0.25))
	sphere3.material.refractiveIndex = 2.5

	ray := NewRay(NewPoint(0, 0, -4), New3Vector(0, 0, 1))

	intersections := NewIntersections(
		&Intersection{sphere1, 2.0},
		&Intersection{sphere2, 2.75},
		&Intersection{sphere3, 3.25},
		&Intersection{sphere2, 4.75},
		&Intersection{sphere3, 5.25},
		&Intersection{sphere1, 6.0},
	)

	precomputed1 := PrepareComputations(intersections.At(0), ray, intersections)
	if precomputed1.RefractiveIndexEnter != 1.0 || precomputed1.RefractiveIndexExit != 1.5 {
		t.Errorf("Precomputed refractive indices not working")
	}

	precomputed2 := PrepareComputations(intersections.At(1), ray, intersections)
	if precomputed2.RefractiveIndexEnter != 1.5 || precomputed2.RefractiveIndexExit != 2.0 {
		t.Errorf("Precomputed refractive indices not working")
	}

	precomputed3 := PrepareComputations(intersections.At(2), ray, intersections)
	if precomputed3.RefractiveIndexEnter != 2.0 || precomputed3.RefractiveIndexExit != 2.5 {
		t.Errorf("Precomputed refractive indices not working")
	}

	precomputed4 := PrepareComputations(intersections.At(3), ray, intersections)
	if precomputed4.RefractiveIndexEnter != 2.5 || precomputed4.RefractiveIndexExit != 2.5 {
		t.Errorf("Precomputed refractive indices not working")
	}

	precomputed5 := PrepareComputations(intersections.At(4), ray, intersections)
	if precomputed5.RefractiveIndexEnter != 2.5 || precomputed5.RefractiveIndexExit != 1.5 {
		t.Errorf("Precomputed refractive indices not working")
	}

	precomputed6 := PrepareComputations(intersections.At(5), ray, intersections)
	if precomputed6.RefractiveIndexEnter != 1.5 || precomputed6.RefractiveIndexExit != 1.0 {
		t.Errorf("Precomputed refractive indices not working")
	}
}

func TestPrecomputedUnderPoint(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), New3Vector(0, 0, 1))
	shape := NewGlassSphere()

	shape.SetTransform(Translation(0, 0, 1))

	intersections := NewIntersections(&Intersection{shape, 5})

	precomputed := PrepareComputations(intersections.Hit(), ray, intersections)

	if precomputed.UnderPoint.z <= Epsilon/2  {
		t.Errorf("Precomputed under point not working")
	}

	if precomputed.UnderPoint.z <= precomputed.Point.z  {
		t.Errorf("Precomputed under point not working")
	}
}
